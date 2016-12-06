package reloading

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	ct "github.com/theplant/containers"
	cb "github.com/theplant/containers/combinators"
)

type reloadableHandler struct {
	handler http.Handler
	page    ct.Page
}

func (rh *reloadableHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	cs, err := rh.page.Containers(req)
	h := req.Header.Get("Accept")
	if err != nil {
		handleErrorJson(res, req, err)
		return
	}

	if h == "application/x-container-list" {
		writeContainerList(res, req, cs)
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
	for i, oc := range ics {
		cw := cb.Wrap("div", cb.Attrs{"data-container-id": fmt.Sprintf("%d", i)}, oc)
		cs = append(cs, cw)
	}
	cs = append(cs, cb.Wrap("script",
		cb.Attrs{"src": "https://cdnjs.cloudflare.com/ajax/libs/fetch/1.0.0/fetch.min.js"}))
	cs = append(cs, cb.ScriptByString(reloadscript))
	return
}

func ReloadablePageHandler(page ct.Page, layout ct.Layout) http.Handler {
	return &reloadableHandler{handler: ct.PageHandler(&wrapWithIdPage{page}, layout), page: page}
}

func writeContainerList(res http.ResponseWriter, req *http.Request, cs []ct.Container) {
	out := map[int]string{}
	clist := strings.Split(req.URL.Query().Get("c"), ",")
	for _, is := range clist {
		i, err := strconv.Atoi(is)
		if err != nil {
			handleErrorJson(res, req, err)
			return
		}

		r, err := cs[i].Render(req)
		if err != nil {
			handleErrorJson(res, req, err)
			return
		}

		out[i] = r
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
