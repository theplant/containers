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

type reloadableEvent interface {
	EventName() string
}

const reloadableContainersKey = "containers.reloading.reloadableContainers"

func (rh *reloadableHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	h := req.Header.Get("Accept")

	if h == "application/x-container-list" {
		reloadableContainers := map[string]ct.Container{}
		req = req.WithContext(context.WithValue(req.Context(), reloadableContainersKey, reloadableContainers))
		_, err := rh.withIdPage.Containers(req) // will update reloadableContainers
		if err != nil {
			handleErrorJson(res, req, err)
			return
		}

		writeContainerList(res, req, reloadableContainers)
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

	var reloadableContainers map[string]ct.Container
	if val := r.Context().Value(reloadableContainersKey); val != nil {
		reloadableContainers = val.(map[string]ct.Container)
	}

	for i, oc := range ics {
		cw := withId.wrapWithId(oc, fmt.Sprintf("%d", i+1), reloadableContainers)
		cs = append(cs, cw)
	}

	return
}

func (withId *wrapWithIdPage) wrapWithId(pc ct.Container, containerId string, reloadableContainers map[string]ct.Container) (rc ct.Container) {

	cval := reflect.ValueOf(pc)
	for cval.Kind() == reflect.Ptr {
		cval = cval.Elem()
	}

	for i := 0; i < cval.NumField(); i++ {
		if !cval.Field(i).CanInterface() {
			continue
		}
		field := cval.Field(i).Interface()
		switch fieldc := field.(type) {
		case ct.Container:
			childId := fmt.Sprintf("%s.%d", containerId, i+1)
			rfc := withId.wrapWithId(fieldc, childId, reloadableContainers)
			cval.Field(i).Set(reflect.ValueOf(rfc))
		}
	}
	switch pc.(type) {
	case reloadableEvent:
		rc = cb.Wrap("div", cb.Attrs{"data-container-id": containerId}, pc)
		if reloadableContainers != nil {
			reloadableContainers[containerId] = pc
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

func writeContainerList(res http.ResponseWriter, req *http.Request, cs map[string]ct.Container) {
	out := map[string]string{}
	clist := strings.Split(req.URL.Query().Get("c"), ",")
	for _, is := range clist {
		c := cs[is]
		if c == nil {
			continue
		}

		r, err := cs[is].Render(req)
		if err != nil {
			handleErrorJson(res, req, err)
			return
		}

		out[is] = r
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
