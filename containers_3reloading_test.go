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
