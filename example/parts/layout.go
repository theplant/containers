package parts

import (
	"context"
	"net/http"
)

func MainLayout(r *http.Request, ctx context.Context, body string) (html string, err error) {
	html = LayoutTemplate(body)
	return
}
