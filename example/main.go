package main

import (
	"log"
	"net/http"

	"github.com/theplant/containers/example/actions"
	"github.com/theplant/containers/example/pages"
	"github.com/theplant/containers/example/parts"
	"github.com/theplant/containers/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/actions/addToCart", actions.AddToCart)
	mux.Handle("/products", handlers.PageHandler(&pages.ProductPage{}, parts.MainLayout))
	mux.Handle("/", handlers.PageHandler(&pages.HomePage{}, parts.MainLayout))
	log.Fatal(http.ListenAndServe(":9000", mux))
}
