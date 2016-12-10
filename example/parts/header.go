package parts

import (
	"net/http"

	"github.com/theplant/containers/example/actions"
	"github.com/theplant/containers/example/parts/templates"
)

type Header struct {
}

func (h *Header) Render(r *http.Request) (html string, err error) {
	html = templates.Header(actions.CartCount)
	return
}
