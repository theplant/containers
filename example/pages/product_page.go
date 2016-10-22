package pages

import (
	"net/http"

	"github.com/theplant/containers"
	"github.com/theplant/containers/example/parts"
)

type ProductPage struct {
}

func (pp *ProductPage) Containers(r *http.Request) (cs []containers.Container, err error) {
	cs = []containers.Container{parts.Header, parts.Product, parts.Footer}
	return
}
