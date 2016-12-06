package pages

import (
	"net/http"

	ct "github.com/theplant/containers"
	cb "github.com/theplant/containers/combinators"
	"github.com/theplant/containers/example/parts"
	rl "github.com/theplant/containers/reloading"
)

type StructuredPage struct {
}

func (sp *StructuredPage) Containers(r *http.Request) (cs []ct.Container, err error) {
	cs = []ct.Container{
		rl.WithReloadEvent("cart_updated", &parts.Header{}),
		cb.Wrap("div", cb.Attrs{"class": "wrapper collection clearfix", "data-bind": "style: bodyWrapperTransform"},
			cb.Wrap("aside", cb.Attrs{"class": "sidebar"},
				cb.ToContainer(ProductAside),
			),
			cb.ToContainer(parts.Product),
		),
		cb.ToContainer(parts.Footer),
	}
	return
}

func ProductAside(r *http.Request) (html string, err error) {
	html = "<div> sidemenu </div>"
	return
}
