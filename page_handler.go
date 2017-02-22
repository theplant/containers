package containers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

/*
PageHandler combine a `Page` and `Layout` to a normal http handler, then you can mount it a a url that fits.
*/
func PageHandler(page Page, layout Layout) http.Handler {
	return &mainHandler{page, layout}
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	log.Println(err)
	if errh, ok := err.(ErrHandler); ok {
		errh.HandleErr(w, r, err)
		return
	}

	eh, ok := r.Context().Value(internalServerErrorHandlerKey).(ErrHandler)
	if ok {
		eh.HandleErr(w, r, err)
		return
	}

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

	var html = buf.String()
	if mh.layout != nil {
		html, err = mh.layout(req, html)
		if err != nil {
			handleError(res, req, err)
			return
		}
	}

	fmt.Fprintln(res, html)
}
