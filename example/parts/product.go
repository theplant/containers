package parts

import (
	"net/http"

	"github.com/theplant/containers/example/models"
	"github.com/theplant/containers/example/parts/templates"
)

func Product(r *http.Request) (html string, err error) {
	ctx := r.Context()

	p := &models.Product{Name: "iPhone 7"}

	if addToCardError := ctx.Value("events.AddToCart.error"); addToCardError != nil {
		html = addToCardError.(error).Error()
		return
	}

	html = templates.Product(p)
	return
}
