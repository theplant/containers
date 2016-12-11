package containers_test

import (
	"fmt"
	"net/http"
	"strings"

	ct "github.com/theplant/containers"
	cb "github.com/theplant/containers/combinators"
)

func ComplicatedHome(r *http.Request) (cs []ct.Container, err error) {
	// Get key parameter in `Page` and pass it to all containers that needed the parameter.
	segs := strings.Split(r.RequestURI, "/")
	productCode := segs[len(segs)-1]

	cs = []ct.Container{
		&Header{},
		&MainContent{
			ProductCode:      productCode,
			ProductBasicInfo: &ProductBasicInfo{productCode},
			ProductImages: &ProductImages{
				ProductCode:      productCode,
				ProductMainImage: cb.ToContainer(ProductMainImage),
			},
			ProductDescription: &ProductDescription{productCode},
		},
		&Footer{},
	}
	return
}

/*
### Setup a nested container tree

Build a tree of containers, by using container struct field and render the children container inside parent container manually.
*/
func ExampleContainer_2nested() {

	http.Handle("/page2", ct.PageHandler(cb.ToPage(ComplicatedHome), nil))
	//Output:

}

// Container MainContent
type MainContent struct {
	ProductCode        string       // use struct field to get inputs from outside of a container.
	ProductBasicInfo   ct.Container // use struct field to pass in a container that used inside another container.
	ProductImages      ct.Container
	ProductDescription ct.Container
}

func (mc *MainContent) Render(r *http.Request) (html string, err error) {
	var basicInfoHtml, imagesHtml, descriptionHtml string

	// inside the container, call `Render` manually to render inside containers.
	if basicInfoHtml, err = mc.ProductBasicInfo.Render(r); err != nil {
		return
	}

	if imagesHtml, err = mc.ProductImages.Render(r); err != nil {
		return
	}

	if descriptionHtml, err = mc.ProductDescription.Render(r); err != nil {
		return
	}

	html = fmt.Sprintf(`
        <div class="main-content">
            %s
            %s
            %s
        </div>
        `,
		basicInfoHtml,
		imagesHtml,
		descriptionHtml,
	)
	return
}

type ProductBasicInfo struct {
	ProductCode string
}

func (bi *ProductBasicInfo) Render(r *http.Request) (html string, err error) {
	db := getDb()
	p := db.GetProduct(bi.ProductCode)
	html = fmt.Sprintf(`
        <div class="basic-info">
            %s
        </div>
    `,
		p.Name,
	)
	return
}

type ProductImages struct {
	ProductCode      string
	ProductMainImage ct.Container
}

func (pi *ProductImages) Render(r *http.Request) (html string, err error) {
	db := getDb()

	var mainImageHtml string
	if mainImageHtml, err = pi.Render(r); err != nil {
		return
	}

	images := db.GetProductImages(pi.ProductCode)
	html = fmt.Sprintf(`
        <div class="images">
            %s
            %v
        </div>
    `,
		mainImageHtml,
		images,
	)
	return
}

func ProductMainImage(r *http.Request) (html string, err error) {
	html = `<div class="product-main-image>main image</div>`
	return
}

type ProductDescription struct {
	ProductCode string
}

func (pd *ProductDescription) Render(r *http.Request) (html string, err error) {
	db := getDb()
	desc := db.GetProductDescription(pd.ProductCode)
	html = fmt.Sprintf(`
        <div class="description">
            %s
        </div>
    `,
		desc,
	)
	return
}
