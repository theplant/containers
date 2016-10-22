package parts

import "net/http"

func MainLayout(r *http.Request, body string) (html string, err error) {
	html = LayoutTemplate(body)
	return
}
