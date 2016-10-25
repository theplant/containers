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
	Page   Page
	Layout Layout
}

func (mh *mainHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	cs, err := mh.Page.Containers(req)
	buf := bytes.NewBuffer(nil)
	for i, c := range cs {
		r, err := c.Render(req)
		if err != nil {
			handleError(res, req, err)
			return
		}
		buf.WriteString(fmt.Sprintf("<div data-container-id=\"%d\">%s</div>", i, r))
	}

	html, err := mh.Layout(req, buf.String())
	if err != nil {
		handleError(res, req, err)
		return
	}
	fmt.Fprintln(res, html)
}
