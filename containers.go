package containers

import (
	"context"
	"net/http"
)

type Container interface {
	Render(r *http.Request, ctx context.Context) (html string, err error)
}

type Page interface {
	Containers(r *http.Request, ctx context.Context) (cs []Container, err error)
}
