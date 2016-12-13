package parts

import (
	"fmt"
	"math/rand"
	"net/http"
)

type ProductColors struct {
}

func (pc *ProductColors) Render(r *http.Request) (html string, err error) {
	html = fmt.Sprintf(`<div style="border: 5px solid yellow">product colors: %d</div>`, rand.Int())
	return
}
