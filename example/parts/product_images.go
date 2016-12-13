package parts

import (
	"fmt"
	"math/rand"
	"net/http"

	ct "github.com/theplant/containers"
)

type ProductImages struct {
	MainImage ct.Container
}

func (prod *ProductImages) Render(r *http.Request) (html string, err error) {
	var mainImageHtml string
	mainImageHtml, err = prod.MainImage.Render(r)
	html = fmt.Sprintf(
		`
        <div class="images" style="border: 5px solid blue">
        %s
        this has many images %d
        </div>
        `,
		mainImageHtml,
		rand.Int(),
	)
	return
}
