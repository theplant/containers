package containers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type MainHandler struct {
	Page Page
}

func (mh *MainHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	cs, _ := mh.Page.Containers(req)
	h := req.Header.Get("Accept")
	if h == "application/x-container-list" {
		writeContainerList(res, req, cs)
	} else {
		writePage(res, req, cs)
	}
}

func writePage(res http.ResponseWriter, req *http.Request, cs []Container) {
	for i, c := range cs {
		r, err := c(req)
		if err != nil {
			log.Println(err)
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
			log.Println(err)
		} else {
			r, _ := cs[i](req)
			if err != nil {
				log.Println(err)
			} else {
				out[i] = r
			}
		}
	}

	json, err := json.Marshal(out)
	if err != nil {
		log.Println(err)
	}
	res.Write(json)
}
