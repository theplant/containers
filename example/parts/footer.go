package parts

import (
	"context"
	"net/http"
)

func Footer(r *http.Request, ctx context.Context) (html string, err error) {
	html = FooterTemplate()
	return
}
