package templates

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
	"github.com/theplant/containers/example/models"
)

func Product(p *models.Product, productImagesHtml string, productColorsHtml string) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n<div style=\"border: 5px solid green\">\n    <h1>")
	_buffer.WriteString(gorazor.HTMLEscape(p.Name))
	_buffer.WriteString("</h1>\n    <button onclick=\"addToCart(111222)\" class=\"addToCart\">Add To Cart</button>\n    ")
	_buffer.WriteString((productImagesHtml))
	_buffer.WriteString("\n    <div class=\"productColors\">\n    ")
	_buffer.WriteString((productColorsHtml))
	_buffer.WriteString("\n    </div>\n</div>\n<script type=\"text/javascript\">\n\nfunction addToCart(variantId) {\n    const body = new FormData();\n    body.append(\"VariantId\", variantId);\n    fetch(\"/actions/addToCart\", {\n        method: \"POST\",\n        body: body,\n    }).then(function(res){\n        postEvent(\"cart_updated\")\n    }).catch(e => alert(e));\n}\n\n</script>")

	return _buffer.String()
}
