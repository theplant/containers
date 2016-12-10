

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




## Type: Container
``` go
type Container interface {
    Render(r *http.Request) (html string, err error)
}
```
Container is a build block for html pape. it's an isolated block that should contains how to prepare data, and how to render itself to html.

By following container interface to implements parts of your html page. You can utilize package built upon containers interface that implements specific features, like reloading, combinators package.



```go
	package containers_test
	
	import (
	    "net/http"
	
	    ct "github.com/theplant/containers"
	)
	
	type Header struct {
	}
	
	func (h *Header) Render(r *http.Request) (html string, err error) {
	    html = "header"
	    return
	}
	
	type Footer struct {
	}
	
	func (h *Footer) Render(r *http.Request) (html string, err error) {
	    html = "footer"
	    return
	}
	
	type Home struct {
	}
	
	func (h *Home) Containers(r *http.Request) (cs []ct.Container, err error) {
	    cs = []ct.Container{
	        &Header{},
	        &Footer{},
	    }
	    return
	}
	
	func ExampleContainer() {
	
	    http.Handle("/page1", ct.PageHandler(&Home{}, nil))
	    //Output:
	
	}
```







## Type: Container Func
``` go
type ContainerFunc func(r *http.Request) (html string, err error)
```









## Type: Layout
``` go
type Layout func(r *http.Request, body string) (html string, err error)
```









## Type: Page
``` go
type Page interface {
    Containers(r *http.Request) (cs []Container, err error)
}
```









## Type: Page Func
``` go
type PageFunc func(r *http.Request) (cs []Container, err error)
```










