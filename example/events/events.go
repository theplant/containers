package events

type CartUpdated struct {
	CartId string
}

type MenuUpdated struct {
}

type AddToCartEvent struct {
	VariantId string
}
