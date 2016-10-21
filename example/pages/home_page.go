package pages

import (
	"context"
	"net/http"

	"github.com/theplant/containers"
)

type HomePage struct {
}

func (hp *HomePage) Containers(r *http.Request, ctx context.Context) (cs []containers.Container, err error) {

	return
}
