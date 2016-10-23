package parts

import (
	"net/http"

	"github.com/theplant/containers/example/actions"
)

func Header(r *http.Request) (html string, err error) {
	html = HeaderTemplate(actions.CartCount)
	return
}
