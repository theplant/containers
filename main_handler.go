package containers

import (
	"log"
	"net/http"
)

type MainHandler struct {
}

func (mh *MainHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Println("setting up routes please")
	return
}
