package templates

import (
	"bytes"
	"github.com/theplant/containers/fetch"
)

func Layout(values *fetch.Results) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\n<html>\n    ")
	_buffer.WriteString((values.Get("head")))
	_buffer.WriteString("\n  <body>\n      <header>\n        ")
	_buffer.WriteString((values.Get("header")))
	_buffer.WriteString("\n      </header>\n      <menu>\n        ")
	_buffer.WriteString((values.Get("menu")))
	_buffer.WriteString("\n      </menu>\n      <content>\n        ")
	_buffer.WriteString((values.Get("content")))
	_buffer.WriteString("\n      </content>\n      <footer>\n        ")
	_buffer.WriteString((values.Get("footer")))
	_buffer.WriteString("\n      </footer>\n  </body>\n</html>")

	return _buffer.String()
}
