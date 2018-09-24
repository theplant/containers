

Package containers provide better structures for building web applications




* [New Redirect Error](#new-redirect-error)
* [Page Handler](#page-handler)
* [Use Err Handler](#use-err-handler)
* [Type Container](#type-container)
* [Type Container Func](#type-container-func)
  * [Render](#container-func-render)
* [Type Err Handler](#type-err-handler)
* [Type Layout](#type-layout)
* [Type Page](#type-page)
* [Type Page Func](#type-page-func)
  * [Containers](#page-func-containers)




## New Redirect Error
``` go
func NewRedirectError(url string, code int) (err error)
```


## Page Handler
``` go
func PageHandler(page Page, layout Layout) http.Handler
```
PageHandler combine a `Page` and `Layout` to a normal http handler, then you can mount it a a url that fits.



## Use Err Handler
``` go
func UseErrHandler(h http.Handler, errhandler ErrHandler) (handler http.Handler)
```




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
	        ct.ContainerFunc(SimpleContent), // Use ct.ContainerFunc to convert a ContainerFunc to a Container
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

### Setup a nested containers tree

Build a tree of containers, by using container struct field and render the children container inside parent container manually.
Note that the struct field name has to be exported, means uppercase. Or it can't benefit `reloading` package.
```go
	package containers_test
	
	import (
	    "fmt"
	    "net/http"
	    "strings"
	
	    ct "github.com/theplant/containers"
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
	                ProductMainImage: ct.ContainerFunc(ProductMainImage),
	            },
	            ProductDescription: &ProductDescription{productCode},
	        },
	        &Footer{},
	    }
	    return
	}
	
	/*
	### Setup a nested containers tree
	
	Build a tree of containers, by using container struct field and render the children container inside parent container manually.
	Note that the struct field name has to be exported, means uppercase. Or it can't benefit `reloading` package.
	*/
	func ExampleContainer_2nested() {
	
	    http.Handle("/page2", ct.PageHandler(ct.PageFunc(ComplicatedHome), nil))
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

If you want to fetch containers separately. Means you can use ajax to only load those containers without rendering other containers by:

- Wrap with `reloading.WithTags` to tag the containers you want to load partially.
- The containers tree struct field for child container must be exported, for example: `&MainContent{ProductBasicInfo: ct}`, can NOT be `&MainContent{productBasicInfo: ct}`
- Use page handler `reloading.ReloadablePageHandler` to mount your `Page` to routes.
- Pass the http header `Accept` with value `application/x-container-list`
- Pass a query parameter called `containersByTags` to a tag names you setup like: `product_updated, description_updated`

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

The result json is a mapping of DOM element container ids inside html page, and rendered html for you to replace into those DOM element. it will look like this:
```javascript
	{
		"2.1": "(ProductBasicInfo container rendered html escaped in json)",
		"2.3": "(ProductDescription container rendered html escaped in json)"
	}
```

"2.1", "2.3" is the attribute value of `<div data-container-id='2.1'></div>` tag inside your html page. So that you can simply replace the server render html to update those tags in javascript.
```go
	package containers_test
	
	import (
	    "net/http"
	    "strings"
	
	    ct "github.com/theplant/containers"
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
	                ProductMainImage: ct.ContainerFunc(ProductMainImage),
	            },
	            ProductDescription: rl.WithTags("product_updated, description_updated", &ProductDescription{productCode}),
	        },
	        &Footer{},
	    }
	    return
	}
	
	/*
	### Fetch certain containers of a page partially
	
	If you want to fetch containers separately. Means you can use ajax to only load those containers without rendering other containers by:
	
	- Wrap with `reloading.WithTags` to tag the containers you want to load partially.
	- The containers tree struct field for child container must be exported, for example: `&MainContent{ProductBasicInfo: ct}`, can NOT be `&MainContent{productBasicInfo: ct}`
	- Use page handler `reloading.ReloadablePageHandler` to mount your `Page` to routes.
	- Pass the http header `Accept` with value `application/x-container-list`
	- Pass a query parameter called `containersByTags` to a tag names you setup like: `product_updated, description_updated`
	
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
	
	The result json is a mapping of DOM element container ids inside html page, and rendered html for you to replace into those DOM element. it will look like this:
	```javascript
	    {
	        "2.1": "(ProductBasicInfo container rendered html escaped in json)",
	        "2.3": "(ProductDescription container rendered html escaped in json)"
	    }
	```
	
	"2.1", "2.3" is the attribute value of `<div data-container-id='2.1'></div>` tag inside your html page. So that you can simply replace the server render html to update those tags in javascript.
	
	*/
	func ExampleContainer_3reloading() {
	
	    http.Handle("/page3", rl.ReloadablePageHandler(ct.PageFunc(ReloadableHome), nil))
	    //Output:
	
	}
```

```go
	http.Handle("/page4", ct.UseErrHandler(ct.PageHandler(ct.PageFunc(MyCatHome), nil), &errhandler{}))
	//Output:
```







## Type: Container Func
``` go
type ContainerFunc func(r *http.Request) (html string, err error)
```
ContainerFunc is a short cut to build simple container that don't depend outside inputs.
Use `containers.ContainerFunc` to convert a `ContainerFunc` to a `Container`.










### Container Func: Render
``` go
func (f ContainerFunc) Render(r *http.Request) (html string, err error)
```



## Type: Err Handler
``` go
type ErrHandler interface {
    HandleErr(w http.ResponseWriter, r *http.Request, err error)
}
```









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
Use `containers.PageFunc` to convert a `PageFunc` to a `Page`










### Page Func: Containers
``` go
func (f PageFunc) Containers(r *http.Request) (cs []Container, err error)
```




