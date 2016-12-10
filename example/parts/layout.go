package parts

import (
	"net/http"

	"github.com/theplant/containers/example/parts/templates"
)

func MainLayout(r *http.Request, body string) (html string, err error) {
	html = templates.Layout(body)
	return
}
