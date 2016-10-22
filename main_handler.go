package containers

import "net/http"

type MainHandler struct {
	Page Page
}

func (mh *MainHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	cs, _ := mh.Page.Containers(req)
	for _, c := range cs {
		r, _ := c(req)
		res.Write([]byte(r))
	}
	return
}
