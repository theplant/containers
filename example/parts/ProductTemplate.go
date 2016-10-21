package parts

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
	"github.com/theplant/containers"
	"github.com/theplant/containers/example/actions"
	"github.com/theplant/containers/example/events"
	"github.com/theplant/containers/example/models"
)

func ProductTemplate(p *models.Product, addToCartEvent *events.AddToCartEvent) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n<div>\n    <h1>p.Name</h1>\n\n    <form>\n        <input type=\"hidden\" name=\"VariantId\"/>\n        <a href=\"#\" data-event=\"")
	_buffer.WriteString(gorazor.HTMLEscape(containers.ActionOn("click", actions.AddToCart, addToCartEvent)))
	_buffer.WriteString("\">Add To Cart</a>\n    </form> \n\n</div>")

	return _buffer.String()
}
