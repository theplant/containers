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
		cb.ToContainer(parts.Product),
		cb.ToContainer(parts.Footer),
	}
	return
}
