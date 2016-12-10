package templates

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
	"strings"
)

func Footer() string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\n<footer style=\"padding: 20px; background-color: #ddd;\">\nThis is a ")
	_buffer.WriteString(gorazor.HTMLEscape(strings.TrimSpace("Footer")))
	_buffer.WriteString("\n</footer>")

	return _buffer.String()
}
