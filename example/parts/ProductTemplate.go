package parts

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
	"github.com/theplant/containers/example/models"
)

func ProductTemplate(p *models.Product) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n<div>\n    <h1>")
	_buffer.WriteString(gorazor.HTMLEscape(p.Name))
	_buffer.WriteString("</h1>\n\n    <form>\n        <input type=\"hidden\" name=\"VariantId\" value=\"111222\"/>\n        <a href=\"#\" data-container-event=\"cart_updated\" class=\"addToCart\">Add To Cart</a>\n    </form> \n\n</div>\n<script src=\"https://cdnjs.cloudflare.com/ajax/libs/fetch/1.0.0/fetch.min.js\"></script>\n<script type=\"text/javascript\">\n\ndocument.addEventListener(\"click\", postAction);\n\nfunction postAction(e) {\n    console.log(e)\n    const event = e.target.dataset.containerEvent\n    if (event != null) {\n        var varid = document.querySelector(\"input[name=VariantId]\")\n        fetch(\"/actions/addToCart\", {\n            VariantId: varid.value\n        }).then(function(res){\n            postEvent(event)\n        })\n    }\n}\n\n</script>")

	return _buffer.String()
}
