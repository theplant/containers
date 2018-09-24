package containers_test

import (
	"net/http"

	ct "github.com/theplant/containers"
)

// Container Header
type Header struct{}

func (h *Header) Render(r *http.Request) (html string, err error) {
	html = "header"
	return
}

// Container Footer
type Footer struct{}

func (h *Footer) Render(r *http.Request) (html string, err error) {
	html = "footer"
	return
}

// ContainerFunc SimpleContent
func SimpleContent(r *http.Request) (html string, err error) {
	html = "simple content"
	return
}

type Home struct {
}

func (h *Home) Containers(r *http.Request) (cs []ct.Container, err error) {
	cs = []ct.Container{
		&Header{},
		ct.ContainerFunc(SimpleContent), // Use ct.ContainerFunc to convert a ContainerFunc to a Container
		&Footer{},
	}
	return
}

/*
### Simple container setup

Mount a `Page` as a `http.Handler` to a http server.
*/
func ExampleContainer_1simple() {

	http.Handle("/page1", ct.PageHandler(&Home{}, nil))
	//Output:

}
