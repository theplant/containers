package reloading

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	c "github.com/theplant/containers"
)

func ReloadingScript() c.Container {
	return c.ScriptByString(reloadscript)
}

type reloadableHandler struct {
	handler http.Handler
	page    c.Page
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
	page c.Page
}

func (withId *wrapWithIdPage) Containers(r *http.Request) (cs []c.Container, err error) {
	var ics []c.Container
	ics, err = withId.page.Containers(r)
	if err != nil {
		return
	}
	for i, oc := range ics {
		cw := c.Wrap(oc, "div", c.Attrs{"data-container-id": fmt.Sprintf("%d", i)})
		cs = append(cs, cw)
	}
	cs = append(cs, c.Wrap(c.StringContainer(""), "script", c.Attrs{"src": "https://cdnjs.cloudflare.com/ajax/libs/fetch/1.0.0/fetch.min.js"}))
	cs = append(cs, ReloadingScript())
	return
}

func ReloadablePageHandler(page c.Page, layout c.Layout) http.Handler {
	return &reloadableHandler{handler: c.PageHandler(&wrapWithIdPage{page}, layout), page: page}
}

func writeContainerList(res http.ResponseWriter, req *http.Request, cs []c.Container) {
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
