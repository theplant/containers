package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/theplant/containers"
	"github.com/theplant/containers/reloading"
)

func PageHandler(page containers.Page, layout containers.Layout) http.Handler {
	return &mainHandler{page, layout}
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, err)
	return
}

type mainHandler struct {
	Page   containers.Page
	Layout containers.Layout
}

func (mh *mainHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	cs, err := mh.Page.Containers(req)
	h := req.Header.Get("Accept")
	if err != nil {
		handleError(res, req, err)
	} else if h == "application/x-container-list" {
		writeContainerList(res, req, cs)
	} else {
		writePage(res, req, cs, mh.Layout)
	}
}

func writePage(res http.ResponseWriter, req *http.Request, cs []containers.Container, layout containers.Layout) {
	buf := bytes.NewBuffer(nil)
	var needReloadScript bool
	for i, c := range cs {
		r, hasReloadable, err := renderWrapC(c, req)
		if err != nil {
			handleError(res, req, err)
			return
		}

		if hasReloadable {
			needReloadScript = true
		}

		buf.WriteString(fmt.Sprintf("<div data-container-id=\"%d\">%s</div>", i, r))
	}

	if needReloadScript {
		r, err := reloading.ReloadingScript().Render(req)
		if err != nil {
			handleError(res, req, err)
			return
		}

		buf.WriteString(r)
	}

	html, err := layout(req, buf.String())
	if err != nil {
		handleError(res, req, err)
		return
	}
	fmt.Fprintln(res, html)
}

func writeContainerList(res http.ResponseWriter, req *http.Request, cs []containers.Container) {
	out := map[int]string{}
	clist := strings.Split(req.URL.Query().Get("c"), ",")
	for _, is := range clist {
		i, err := strconv.Atoi(is)
		if err != nil {
			handleError(res, req, err)
			return
		}

		r, _, _ := renderWrapC(cs[i], req)
		if err != nil {
			handleError(res, req, err)
			return
		}

		out[i] = r
	}

	json, err := json.Marshal(out)
	if err != nil {
		handleError(res, req, err)
	}
	res.Write(json)
}

func renderWrapC(c containers.Container, r *http.Request) (html string, hasReloadable bool, err error) {
	var rl containers.Reloadable
	if rl, hasReloadable = c.(containers.Reloadable); hasReloadable {
		c = reloading.Reloadable(rl.ReloadEvent(), c)
	}
	html, err = c.Render(r)
	return
}
