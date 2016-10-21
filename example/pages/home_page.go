package pages

import (
	"context"
	"net/http"

	"github.com/theplant/containers"
	"github.com/theplant/containers/example/parts"
)

type HomePage struct {
}

func (hp *HomePage) Containers(r *http.Request, ctx context.Context) (cs []containers.Container, err error) {
	cs = []containers.Container{parts.Header, parts.Product, parts.Footer}

	return
}
