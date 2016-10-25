package pages

import (
	"net/http"

	c "github.com/theplant/containers"
	"github.com/theplant/containers/example/parts"
	rl "github.com/theplant/containers/reloading"
)

type ProductPage struct {
}

func (pp *ProductPage) Containers(r *http.Request) (cs []c.Container, err error) {
	cs = []c.Container{
		rl.WithReloadEvent("cart_updated", &parts.Header{}),
		c.ContainerFunc(parts.Product),
		c.ContainerFunc(parts.Footer),
	}
	return
}
