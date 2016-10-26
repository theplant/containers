package actions

import (
	"log"
	"net/http"
)

var CartCount int

func AddToCart(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	var variantId = r.FormValue("VariantId")
	log.Println("AddToCart called: ", variantId)
	if len(variantId) == 0 {
		// err = errors.New("Need to provide variant id")
		// return
	}

	// Do database operation for add_to_cart

	CartCount++

	return
}
