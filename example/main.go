package main

import (
	"log"
	"net/http"

	"github.com/theplant/containers"
	"github.com/theplant/containers/example/pages"
	"github.com/theplant/containers/example/parts"
)

func main() {
	containers.GET("/products/:code", &pages.ProductPage{}, parts.MainLayout)
	containers.GET("/", &pages.HomePage{}, parts.MainLayout)

	http.Handle("/", &containers.MainHandler{})
	log.Fatal(http.ListenAndServe(":9000", nil))
}
