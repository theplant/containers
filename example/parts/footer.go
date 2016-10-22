package parts

import "net/http"

func Footer(r *http.Request) (html string, err error) {
	html = FooterTemplate()
	return
}
