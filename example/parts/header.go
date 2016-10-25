package parts

import (
	"net/http"

	"github.com/theplant/containers/example/actions"
)

type Header struct {
}

func (h *Header) Render(r *http.Request) (html string, err error) {
	html = HeaderTemplate(actions.CartCount)
	return
}

func (h *Header) ReloadEvent() string {
	return "cart_updated"
}
