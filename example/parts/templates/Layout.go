package templates

import (
	"bytes"
)

func Layout(body string) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n<html>\n    <head>Containers Example</head>\n  <body>\n      ")
	_buffer.WriteString((body))
	_buffer.WriteString("\n  </body>\n</html>")

	return _buffer.String()
}
