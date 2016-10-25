package reloading

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/theplant/containers"
)

func ReloadingScript() containers.Container {
	return containers.Script("reload.js")
}

type reloadableHandler struct {
	handler http.Handler
	page    containers.Page
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

func ReloadablePageHandler(page containers.Page, layout containers.Layout) http.Handler {
	return &reloadableHandler{handler: containers.PageHandler(page, layout), page: page}
}

func writeContainerList(res http.ResponseWriter, req *http.Request, cs []containers.Container) {
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
