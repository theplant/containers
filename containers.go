package containers

import "net/http"

type Container interface {
	Render(r *http.Request) (html string, err error)
}

func ContainerFunc(f func(r *http.Request) (html string, err error)) Container {
	return containerFunc{f}
}

type containerFunc struct {
	cf func(r *http.Request) (html string, err error)
}

func (f containerFunc) Render(r *http.Request) (html string, err error) {
	return f.cf(r)
}

type Layout func(r *http.Request, body string) (html string, err error)

type Action func(r *http.Request) (redirectUrl string, events []Event, err error)

type Event interface {
}

type Page interface {
	Containers(r *http.Request) (cs []Container, err error)
}

func Handler(page Page, layout Layout) http.Handler {

	return &MainHandler{page, layout}
}

func ReloadContainerOn(c Container, events ...string) {
	return
}
