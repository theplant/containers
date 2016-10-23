package main

import (
	"log"
	"net/http"

	"github.com/theplant/containers"
	"github.com/theplant/containers/example/pages"
)

func main() {
	mux := http.NewServeMux()
	pages.AddRoutes(mux)
	mux.Handle("/", &containers.MainHandler{&pages.HomePage{}})
	log.Fatal(http.ListenAndServe(":9000", mux))
}
