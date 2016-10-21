package parts

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
	"strings"
)

func FooterTemplate() string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\n<footer>Hello ")
	_buffer.WriteString(gorazor.HTMLEscape(strings.TrimSpace("Footer")))
	_buffer.WriteString("</footer>")

	return _buffer.String()
}
