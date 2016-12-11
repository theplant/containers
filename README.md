

Package containers provide better structures for building web applications




* [Page Handler](#page-handler)
* [Type Container](#type-container)
* [Type Container Func](#type-container-func)
* [Type Layout](#type-layout)
* [Type Page](#type-page)
* [Type Page Func](#type-page-func)




## Page Handler
``` go
func PageHandler(page Page, layout Layout) http.Handler
```
PageHandler combine a `Page` and `Layout` to a normal http handler, then you can mount it a a url that fits.





## Type: Container
``` go
type Container interface {
    Render(r *http.Request) (html string, err error)
}
```
Container is a build block for html pape. it's an isolated block that should contains how to prepare data, and how to render itself to html.

By following container interface to implements parts of your html page. You can utilize package built upon containers interface that implements specific features, like reloading, combinators package.



### Simple container setup

Mount a `Page` as a `http.Handler` to a http server.
```go
	package containers_test
	
	import (
	    "net/http"
	
	    ct "github.com/theplant/containers"
	    cb "github.com/theplant/containers/combinators"
	)
	
	// Container Header
	type Header struct{}
	
	func (h *Header) Render(r *http.Request) (html string, err error) {
	    html = "header"
	    return
	}
	
	// Container Footer
	type Footer struct{}
	
	func (h *Footer) Render(r *http.Request) (html string, err error) {
	    html = "footer"
	    return
	}
	
	// ContainerFunc SimpleContent
	func SimpleContent(r *http.Request) (html string, err error) {
	    html = "simple content"
	    return
	}
	
	type Home struct {
	}
	
	func (h *Home) Containers(r *http.Request) (cs []ct.Container, err error) {
	    cs = []ct.Container{
	        &Header{},
	        cb.ToContainer(SimpleContent), // Use combinators.ToContainer to convert a ContainerFunc to a Container
	        &Footer{},
	    }
	    return
	}
	
	/*
	### Simple container setup
	
	Mount a `Page` as a `http.Handler` to a http server.
	*/
	func ExampleContainer_1simple() {
	
	    http.Handle("/page1", ct.PageHandler(&Home{}, nil))
	    //Output:
	
	}
```

### Setup a nested container tree

Build a tree of containers, by using container struct field and render the children container inside parent container manually.
```go
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
```

### Fetch certain containers of a page partially

If you wrap your container inside your page with `reloading.WithTags`, then those containers are separate fetchable.
means you can use ajax to only load those containers without rendering other containers by:

- pass the http header `Accept` with value `application/x-container-list`
- pass a query parameter called `containersByTags` to a tag names you setup like: `product_updated, description_updated`

In this page example:

```javascript

    const url = "/page3?containersByTags=product_updated"
    const reqData = {
        headers: {
            'Accept': "application/x-container-list"
        }
    }
    fetch(url, reqData).then(r => r.json())
```

The result json is a mapping of DOM element container ids inside html page, and rendered html for you to replace into those DOM element.
```go
	package containers_test
	
	import (
	    "net/http"
	    "strings"
	
	    ct "github.com/theplant/containers"
	    cb "github.com/theplant/containers/combinators"
	    rl "github.com/theplant/containers/reloading"
	)
	
	func ReloadableHome(r *http.Request) (cs []ct.Container, err error) {
	    // Get key parameter in `Page` and pass it to all containers that needed the parameter.
	    segs := strings.Split(r.RequestURI, "/")
	    productCode := segs[len(segs)-1]
	
	    cs = []ct.Container{
	        &Header{},
	        &MainContent{
	            ProductCode:      productCode,
	            ProductBasicInfo: rl.WithTags("product_updated", &ProductBasicInfo{productCode}),
	            ProductImages: &ProductImages{
	                ProductCode:      productCode,
	                ProductMainImage: cb.ToContainer(ProductMainImage),
	            },
	            ProductDescription: rl.WithTags("product_updated, description_updated", &ProductDescription{productCode}),
	        },
	        &Footer{},
	    }
	    return
	}
	
	/*
	### Fetch certain containers of a page partially
	
	If you wrap your container inside your page with `reloading.WithTags`, then those containers are separate fetchable.
	means you can use ajax to only load those containers without rendering other containers by:
	
	- pass the http header `Accept` with value `application/x-container-list`
	- pass a query parameter called `containersByTags` to a tag names you setup like: `product_updated, description_updated`
	
	In this page example:
	
	```javascript
	
	    const url = "/page3?containersByTags=product_updated"
	    const reqData = {
	        headers: {
	            'Accept': "application/x-container-list"
	        }
	    }
	    fetch(url, reqData).then(r => r.json())
	```
	
	The result json is a mapping of DOM element container ids inside html page, and rendered html for you to replace into those DOM element.
	*/
	func ExampleContainer_3reloading() {
	
	    http.Handle("/page3", ct.PageHandler(cb.ToPage(ReloadableHome), nil))
	    //Output:
	
	}
```







## Type: Container Func
``` go
type ContainerFunc func(r *http.Request) (html string, err error)
```
ContainerFunc is a short cut to build simple container that don't depend outside inputs.
Use `combinators.ToContainer` to convert a `ContainerFunc` to a `Container`.










## Type: Layout
``` go
type Layout func(r *http.Request, body string) (html string, err error)
```
Layout is like a `Container`, but takes another parameter body. use `containers.PageHandler` to combine a `Page` and a `Layout`, and mount to a url.










## Type: Page
``` go
type Page interface {
    Containers(r *http.Request) (cs []Container, err error)
}
```
Page is anything that can return a list of `Containers`, use `containers.PageHandler` to convert a `Page` to a `http.Handler`, So that you can mount it to a url.










## Type: Page Func
``` go
type PageFunc func(r *http.Request) (cs []Container, err error)
```
PageFunc is a short cut to build simple page that don't depend outside inputs.
Use `combinators.ToPage` to convert a `PageFunc` to a `Page`











