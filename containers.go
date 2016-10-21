package containers

import (
	"context"
	"net/http"
)

type Container func(r *http.Request, ctx context.Context) (html string, err error)

type Layout func(r *http.Request, ctx context.Context, body string) (html string, err error)

type Page interface {
	Containers(r *http.Request, ctx context.Context) (cs []Container, err error)
}

func GET(relativePath string, page Page, layout Layout) {
	return
}
