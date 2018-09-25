package example

import (
	"net/http"

	ct "github.com/theplant/containers"
	"github.com/theplant/containers/fetch"
	"github.com/theplant/containers/fetch/example/templates"
)

var Page1 ct.Page = &fetch.Group{
	BeforeFetch: func(r *http.Request, child *http.Request) error {
		child.Header.Set("My-Header1", "3")

		values := r.URL.Query()
		values.Add("My-Param1", "4")
		child.URL.RawQuery = values.Encode()
		return nil
	},
	CMap: map[string]ct.Container{
		"head":   &fetch.Fetcher{URL: "www.qq.com"},
		"header": &fetch.Fetcher{URL: "www.baidu.com"},
		"menu":   &fetch.Fetcher{URL: "www.baidu.com", Async: true, TTL: 5},
		"content": &fetch.Fetcher{URL: "www.baidu.com", BeforeFetch: func(r *http.Request, child *http.Request) error {
			child.Header.Set("My-Header", "1")

			values := r.URL.Query()
			values.Add("My-Param", "2")
			child.URL.RawQuery = values.Encode()
			return nil
		}},
		"footer": &fetch.Fetcher{URL: "www.163.com"},
	},
	LayoutFunc: templates.Layout,
}
