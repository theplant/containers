package actions

import (
	"context"
	"errors"
	"net/http"

	"github.com/theplant/containers"
	"github.com/theplant/containers/example/events"
)

func AddToCart(r *http.Request, ctx context.Context) (redirectUrl string, evts []containers.Event, err error) {
	var addToCartEvent *events.AddToCartEvent
	if e := ctx.Value("events.AddToCartEvent"); e != nil {
		addToCartEvent = e.(*events.AddToCartEvent)
	}
	if addToCartEvent == nil {
		err = errors.New("Need to provide variant id")
		return
	}

	evts = []containers.Event{&events.CartUpdated{CartId: "111"}}
	return
}
