package main

import (
	"log"
	"net/http"

	"github.com/theplant/containers/example/actions"
	"github.com/theplant/containers/example/pages"
	"github.com/theplant/containers/example/parts"
	"github.com/theplant/containers/reloading"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/actions/addToCart", actions.AddToCart)
	mux.Handle("/products", reloading.ReloadablePageHandler(&pages.ProductPage{}, parts.MainLayout))
	mux.Handle("/structured_page", reloading.ReloadablePageHandler(&pages.StructuredPage{}, parts.MainLayout))
	mux.Handle("/", reloading.ReloadablePageHandler(&pages.HomePage{}, parts.MainLayout))
	log.Fatal(http.ListenAndServe(":9000", mux))
}
