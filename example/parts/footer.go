package parts

import (
	"net/http"

	"github.com/theplant/containers/example/parts/templates"
)

func Footer(r *http.Request) (html string, err error) {
	html = templates.Footer()
	return
}
