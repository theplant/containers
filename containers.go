package containers

import "net/http"

type Container interface {
	Render(r *http.Request) (html string, err error)
}

type Page interface {
	Containers(r *http.Request) (cs []Container, err error)
}

type Layout func(r *http.Request, body string) (html string, err error)
