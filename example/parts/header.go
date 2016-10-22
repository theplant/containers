package parts

import (
	"net/http"

	"github.com/theplant/containers"
	"github.com/theplant/containers/example/models"
)

func Header(r *http.Request) (html string, err error) {
	html = HeaderTemplate(&models.Product{Name: "Felix"})
	return
}

func init() {
	containers.ReloadContainerOn(Header, "events.CartUpdated", "events.MenuUpdated")
}
