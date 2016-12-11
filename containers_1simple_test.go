package containers_test

import (
	"net/http"

	ct "github.com/theplant/containers"
	cb "github.com/theplant/containers/combinators"
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
		cb.ToContainer(SimpleContent), // Use combinators.ToContainer to convert a ContainerFunc to a Container
		&Footer{},
	}
	return
}

/*
Mount a `Page` as a `http.Handler` to a http server.
*/
func ExampleContainer_1simple() {

	http.Handle("/page1", ct.PageHandler(&Home{}, nil))
	//Output:

}
