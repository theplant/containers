package parts

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func LayoutTemplate(body string) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n<html>\n    <head>Containers Example</head>\n  <body>\n      ")
	_buffer.WriteString(gorazor.HTMLEscape(body))
	_buffer.WriteString("\n  </body>\n</html>")

	return _buffer.String()
}
