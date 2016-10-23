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

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, err)
	return
}

type MainHandler struct {
	Page   Page
	Layout Layout
}

func (mh *MainHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
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

func writePage(res http.ResponseWriter, req *http.Request, cs []Container, layout Layout) {
	buf := bytes.NewBuffer(nil)
	for i, c := range cs {
		r, err := c.Render(req)
		if err != nil {
			handleError(res, req, err)
			return
		}

		buf.WriteString(fmt.Sprintf("<div data-container-id=\"%d\">%s</div>", i, r))
	}

	html, err := layout(req, buf.String())
	if err != nil {
		handleError(res, req, err)
		return
	}
	fmt.Fprintln(res, html)
}

func writeContainerList(res http.ResponseWriter, req *http.Request, cs []Container) {
	out := map[int]string{}
	clist := strings.Split(req.URL.Query().Get("c"), ",")
	for _, is := range clist {
		i, err := strconv.Atoi(is)
		if err != nil {
			handleError(res, req, err)
		} else {
			r, _ := cs[i].Render(req)
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
