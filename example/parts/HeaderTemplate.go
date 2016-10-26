package parts

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func HeaderTemplate(cartCount int) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n<header style=\"padding: 20px; background-color: #ddd;\">\n  This is a header, cart(")
	_buffer.WriteString(gorazor.HTMLEscape(cartCount))
	_buffer.WriteString(")\n</header>")

	return _buffer.String()
}
