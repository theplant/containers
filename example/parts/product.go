package parts

import (
	"net/http"

	ct "github.com/theplant/containers"
	"github.com/theplant/containers/example/models"
	"github.com/theplant/containers/example/parts/templates"
)

type Product struct {
	ProductImages ct.Container
	ProductColors ct.Container
}

func (prod *Product) Render(r *http.Request) (html string, err error) {
	// ctx := r.Context()

	var productImagesHtml, productColorsHtml string

	if prod.ProductImages != nil {
		productImagesHtml, err = prod.ProductImages.Render(r)
		if err != nil {
			return
		}
	}

	if prod.ProductColors != nil {
		productColorsHtml, err = prod.ProductColors.Render(r)
		if err != nil {
			return
		}
	}

	p := &models.Product{Name: "iPhone 7"}

	html = templates.Product(p, productImagesHtml, productColorsHtml)
	return
}
