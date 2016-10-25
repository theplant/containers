package pages

import (
	"net/http"

	c "github.com/theplant/containers"
	"github.com/theplant/containers/example/parts"
)

type ProductPage struct {
}

func (pp *ProductPage) Containers(r *http.Request) (cs []c.Container, err error) {
	cs = []c.Container{&parts.Header{}, c.ContainerFunc(parts.Product), c.ContainerFunc(parts.Footer)}
	return
}
