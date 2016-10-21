package parts

import (
	"context"
	"net/http"

	"github.com/theplant/containers/example/events"
	"github.com/theplant/containers/example/models"
)

func Product(r *http.Request, ctx context.Context) (html string, err error) {
	p := &models.Product{Name: "Felix"}

	var addToCartEvent *events.AddToCartEvent
	if e := ctx.Value("events.AddToCartEvent"); e != nil {
		addToCartEvent = e.(*events.AddToCartEvent)
	}

	if addToCardError := ctx.Value("events.AddToCart.error"); addToCardError != nil {
		html = addToCardError.(error).Error()
		return
	}

	html = ProductTemplate(p, addToCartEvent)
	return
}
