package containers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func PageHandler(page Page, layout Layout) http.Handler {
	return &mainHandler{page, layout}
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, err)
	return
}

type mainHandler struct {
	page   Page
	layout Layout
}

func (mh *mainHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	cs, err := mh.page.Containers(req)
	buf := bytes.NewBuffer(nil)
	for _, c := range cs {
		r, err := c.Render(req)
		if err != nil {
			handleError(res, req, err)
			return
		}
		buf.WriteString(r)
	}

	html, err := mh.layout(req, buf.String())
	if err != nil {
		handleError(res, req, err)
		return
	}
	fmt.Fprintln(res, html)
}
