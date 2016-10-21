package parts

import (
	"context"
	"net/http"

	"github.com/theplant/containers/example/models"
)

func Header(r *http.Request, ctx context.Context) (html string, err error) {
	html = HeaderTemplate(&models.Product{Name: "Felix"})
	return
}
