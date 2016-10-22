package main

import (
	"log"
	"net/http"

	"github.com/theplant/containers/example/pages"
)

func main() {
	mux := http.NewServeMux()
	pages.AddRoutes(mux)
	log.Fatal(http.ListenAndServe(":9000", mux))
}
