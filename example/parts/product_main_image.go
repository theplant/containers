package parts

import (
	"fmt"
	"math/rand"
	"net/http"
)

type ProductMainImage struct {
}

func (pmi *ProductMainImage) Render(r *http.Request) (html string, err error) {
	html = fmt.Sprintf(`<div class="main-image" style="border: 5px solid red">product main image tag: %d</div>`, rand.Int())
	return
}
