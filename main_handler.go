package containers

import "net/http"

type MainHandler struct {
	Page Page
}

func (mh *MainHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	cs, _ := mh.Page.Containers(req, req.Context())
	for _, c := range cs {
		r, _ := c(req, req.Context())
		res.Write([]byte(r))
	}
	return
}
