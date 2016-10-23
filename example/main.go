package main

import (
	"log"
	"net/http"

	"github.com/theplant/containers"
	"github.com/theplant/containers/example/actions"
	"github.com/theplant/containers/example/pages"
	"github.com/theplant/containers/example/parts"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/actions/addToCart", actions.AddToCart)
	mux.Handle("/products", containers.Handler(&pages.ProductPage{}, parts.MainLayout))
	mux.Handle("/", containers.Handler(&pages.HomePage{}, parts.MainLayout))
	log.Fatal(http.ListenAndServe(":9000", mux))
}
