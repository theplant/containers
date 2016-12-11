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
			ProductBasicInfo: rl.WithReloadEvent("product_updated", &ProductBasicInfo{productCode}),
			ProductImages: &ProductImages{
				ProductCode:      productCode,
				ProductMainImage: cb.ToContainer(ProductMainImage),
			},
			ProductDescription: rl.WithReloadEvent("product_updated, description_updated", &ProductDescription{productCode}),
		},
		&Footer{},
	}
	return
}

/*
### Reloading certain containers of a page partially

If you wrap your container inside your page with `reloading.WithReloadEvents`, then those containers are reloadable.
means you can use ajax to only load those containers without rendering other containers by:

- pass the http header `Accept` with value `application/x-container-list`
- pass a query parameter called `reloadingContainerEvents` to a event name you setup like: `product_updated`

In this page example:

```javascript

    const url = "/page3?reloadingContainerEvents=product_updated"
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