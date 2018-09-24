package reloading

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	ct "github.com/theplant/containers"
	cb "github.com/theplant/containers/combinators"
)

type reloadableHandler struct {
	handler    http.Handler
	withIdPage *wrapWithIdPage
}

type taggedContainer interface {
	TagNames() string
}

const containersByTagKey = "containers.reloading.containersByTag"
const ContainersByTagsRequestParamName = "containersByTags"

type containerWithId struct {
	containerId string
	container   ct.Container
}
type tagNameContainersWithId struct {
	tagName        string
	containersById []containerWithId
}
type containersByTag struct {
	lists []*tagNameContainersWithId
}

func (cbt *containersByTag) Put(tagNames string, cId string, c ct.Container) {
	names := pureTagNames(tagNames)
	for _, name := range names {
		var added bool
		for _, ic := range cbt.lists {
			if ic.tagName == name {
				ic.containersById = append(ic.containersById, containerWithId{cId, c})
				added = true
				break
			}
		}
		if !added {
			cbt.lists = append(cbt.lists, &tagNameContainersWithId{name, []containerWithId{
				{cId, c},
			}})
		}
	}
}

func (cbt *containersByTag) Get(tagNames string) (cs []containerWithId) {
	names := pureTagNames(tagNames)
	for _, name := range names {
		for _, ic := range cbt.lists {
			if ic.tagName == name {
				cs = append(cs, ic.containersById...)
			}
		}
	}
	return
}

func pureTagNames(names string) (r []string) {
	ns := strings.Split(names, ",")
	for _, name := range ns {
		name = strings.TrimSpace(name)
		if len(name) == 0 {
			continue
		}
		r = append(r, name)
	}
	return
}

func (rh *reloadableHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	h := req.Header.Get("Accept")

	if h == "application/x-container-list" {
		cbt := &containersByTag{}
		req = req.WithContext(context.WithValue(req.Context(), containersByTagKey, cbt))
		_, err := rh.withIdPage.Containers(req) // will update reloadableContainers
		if err != nil {
			handleErrorJson(res, req, err)
			return
		}

		writeContainerList(res, req, cbt)
		return
	}

	rh.handler.ServeHTTP(res, req)
}

type wrapWithIdPage struct {
	page ct.Page
}

func (withId *wrapWithIdPage) Containers(r *http.Request) (cs []ct.Container, err error) {
	var ics []ct.Container
	ics, err = withId.page.Containers(r)
	if err != nil {
		return
	}

	var cbt *containersByTag
	if val := r.Context().Value(containersByTagKey); val != nil {
		cbt = val.(*containersByTag)
	}

	for i, oc := range ics {
		cw := withId.wrapWithId(oc, fmt.Sprintf("%d", i+1), cbt)
		cs = append(cs, cw)
	}

	return
}

func (withId *wrapWithIdPage) wrapWithId(pc ct.Container, containerId string, cbt *containersByTag) (rc ct.Container) {

	cval := reflect.ValueOf(pc)
	for cval.Kind() == reflect.Ptr {
		cval = cval.Elem()
	}

	if cval.Kind() != reflect.Func {
		for i := 0; i < cval.NumField(); i++ {
			if !cval.Field(i).CanInterface() {
				continue
			}
			field := cval.Field(i).Interface()
			switch fieldc := field.(type) {
			case ct.Container:
				childId := fmt.Sprintf("%s.%d", containerId, i+1)
				rfc := withId.wrapWithId(fieldc, childId, cbt)
				cval.Field(i).Set(reflect.ValueOf(rfc))
			}
		}
	}

	switch tagc := pc.(type) {
	case taggedContainer:
		rc = cb.Wrap("div", cb.Attrs{"data-container-id": containerId}, pc)
		if cbt != nil {
			cbt.Put(tagc.TagNames(), containerId, pc)
		}
	default:
		rc = pc
	}
	return
}

func ReloadablePageHandler(page ct.Page, layout ct.Layout) http.Handler {
	withIdPage := &wrapWithIdPage{page: page}
	return &reloadableHandler{handler: ct.PageHandler(withIdPage, layout), withIdPage: withIdPage}
}

func writeContainerList(res http.ResponseWriter, req *http.Request, cbt *containersByTag) {
	out := map[string]string{}
	cs := cbt.Get(req.URL.Query().Get(ContainersByTagsRequestParamName))
	for _, is := range cs {
		r, err := is.container.Render(req)
		if err != nil {
			handleErrorJson(res, req, err)
			return
		}
		out[is.containerId] = r
	}

	json, err := json.Marshal(out)
	if err != nil {
		handleErrorJson(res, req, err)
	}
	res.Write(json)
}

func handleErrorJson(w http.ResponseWriter, r *http.Request, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, err)
	return
}
