package containers

import "net/http"

type Container func(r *http.Request) (html string, err error)

type Layout func(r *http.Request, body string) (html string, err error)

type Action func(r *http.Request) (redirectUrl string, events []Event, err error)

type Event interface {
}

type Page interface {
	Containers(r *http.Request) (cs []Container, err error)
}

func GET(relativePath string, page Page, layout Layout) {
	return
}

func ReloadContainerOn(c Container, events ...string) {
	return
}
