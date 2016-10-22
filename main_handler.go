package containers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

type HandleFuncMux interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}

func mainHandleFunc(page Page, layout Layout) (r func(http.ResponseWriter, *http.Request)) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var cs []Container
		cs, err = page.Containers(r)
		if err != nil {
			handleError(w, r, err)
			return
		}

		var html string
		buf := bytes.NewBuffer(nil)

		for _, c := range cs {
			html, err = c(r)
			if err != nil {
				handleError(w, r, err)
				return
			}
			buf.WriteString(html)
		}
		html, err = layout(r, buf.String())
		if err != nil {
			handleError(w, r, err)
			return
		}

		_, err = fmt.Fprintln(w, html)
		if err != nil {
			handleError(w, r, err)
		}
		return
	}
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, err)
	return
}
