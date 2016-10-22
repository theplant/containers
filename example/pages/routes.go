package pages

import (
	"github.com/theplant/containers"
	"github.com/theplant/containers/example/parts"
)

func AddRoutes(mux containers.HandleFuncMux) {
	containers.GET(mux, "/products", &ProductPage{}, parts.MainLayout)
	containers.GET(mux, "/", &HomePage{}, parts.MainLayout)
}
