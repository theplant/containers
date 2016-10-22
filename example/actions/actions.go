package actions

import (
	"errors"
	"net/http"

	"github.com/theplant/containers"
	"github.com/theplant/containers/example/events"
)

func AddToCart(r *http.Request) (redirectUrl string, evts []containers.Event, err error) {
	ctx := r.Context()
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
