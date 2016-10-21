package main

import (
	"log"
	"net/http"

	"github.com/theplant/containers"
	"github.com/theplant/containers/example/pages"
	"github.com/theplant/containers/example/parts"
)

func main() {
	containers.GET("/products/:name", &pages.ProductPage{}, parts.MainLayout)

	http.Handle("/", &containers.MainHandler{})
	log.Fatal(http.ListenAndServe(":9000", nil))
}
