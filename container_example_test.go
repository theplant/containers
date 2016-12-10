package containers_test

import (
	"net/http"

	ct "github.com/theplant/containers"
)

type Header struct {
}

func (h *Header) Render(r *http.Request) (html string, err error) {
	html = "header"
	return
}

type Footer struct {
}

func (h *Footer) Render(r *http.Request) (html string, err error) {
	html = "footer"
	return
}

type Home struct {
}

func (h *Home) Containers(r *http.Request) (cs []ct.Container, err error) {
	cs = []ct.Container{
		&Header{},
		&Footer{},
	}
	return
}

func ExampleContainer() {

	http.Handle("/page1", ct.PageHandler(&Home{}, nil))
	//Output:

}
