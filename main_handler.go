package containers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
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
			html, err = c.Content(r)
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

type MainHandler struct {
	Page Page
}

func (mh *MainHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	cs, err := mh.Page.Containers(req)
	h := req.Header.Get("Accept")
	if err != nil {
		handleError(res, req, err)
	} else if h == "application/x-container-list" {
		writeContainerList(res, req, cs)
	} else {
		writePage(res, req, cs)
	}
}

func writePage(res http.ResponseWriter, req *http.Request, cs []Container) {
	for i, c := range cs {
		r, err := c.Content(req)
		if err != nil {
			handleError(res, req, err)
		} else {
			res.Write([]byte(fmt.Sprintf("<div data-container-id=\"%d\">%s</div>", i, r)))
		}
	}
}

func writeContainerList(res http.ResponseWriter, req *http.Request, cs []Container) {
	out := map[int]string{}
	clist := strings.Split(req.URL.Query().Get("c"), ",")
	for _, is := range clist {
		i, err := strconv.Atoi(is)
		if err != nil {
			handleError(res, req, err)
		} else {
			r, _ := cs[i].Content(req)
			if err != nil {
				handleError(res, req, err)
			} else {
				out[i] = r
			}
		}
	}

	json, err := json.Marshal(out)
	if err != nil {
		handleError(res, req, err)
	}
	res.Write(json)
}
