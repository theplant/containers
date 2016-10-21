package parts

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
	"github.com/theplant/containers/example/models"
	"strings"
)

func HeaderTemplate(p *models.Product) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n<header>\n  <h1>Hello ")
	_buffer.WriteString(gorazor.HTMLEscape(strings.TrimSpace(p.Name)))
	_buffer.WriteString("</h1>\n</header>")

	return _buffer.String()
}
