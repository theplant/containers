package pages

import (
	"net/http"

	ct "github.com/theplant/containers"
	cb "github.com/theplant/containers/combinators"
	"github.com/theplant/containers/example/parts"
	rl "github.com/theplant/containers/reloading"
)

type ProductPage struct {
}

func (pp *ProductPage) Containers(r *http.Request) (cs []ct.Container, err error) {
	cs = []ct.Container{
		rl.WithReloadEvent("cart_updated", &parts.Header{}),
		&parts.Product{
			ProductColors: rl.WithReloadEvent("cart_updated", &parts.ProductColors{}),
			ProductImages: &parts.ProductImages{
				MainImage: rl.WithReloadEvent("cart_updated", &parts.ProductMainImage{}),
			},
		},
		cb.ToContainer(parts.Footer),
		cb.Wrap("script", cb.Attrs{"src": "https://cdnjs.cloudflare.com/ajax/libs/fetch/1.0.0/fetch.min.js"}),
		cb.ScriptByString(rl.ReloadScript),
	}
	return
}
