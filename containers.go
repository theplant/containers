package containers

import "net/http"

type Container interface {
	Content(r *http.Request) (html string, err error)
}

func ContainerFunc(f func(r *http.Request) (html string, err error)) Container {
	return containerFunc{f}
}

type containerFunc struct {
	cf func(r *http.Request) (html string, err error)
}

func (f containerFunc) Content(r *http.Request) (html string, err error) {
	return f.cf(r)
}

type Layout func(r *http.Request, body string) (html string, err error)

type Action func(r *http.Request) (redirectUrl string, events []Event, err error)

type Event interface {
}

type Page interface {
	Containers(r *http.Request) (cs []Container, err error)
}

func GET(mux HandleFuncMux, pattern string, page Page, layout Layout) {
	mux.HandleFunc(pattern, mainHandleFunc(page, layout))
	return
}

func ReloadContainerOn(c Container, events ...string) {
	return
}
