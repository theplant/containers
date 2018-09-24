package example

import (
	"net/http"

	ct "github.com/theplant/containers"
	"github.com/theplant/containers/fetch"
	"github.com/theplant/containers/fetch/example/templates"
)

func Page1(r *http.Request) []ct.Container {
	return []ct.Container{
		&fetch.Group{
			BeforeFetch: func(r *http.Request, child *http.Request) error {
				child.Header.Set("My-Header1", "3")

				values := r.URL.Query()
				values.Add("My-Param1", "4")
				child.URL.RawQuery = values.Encode()
				return nil
			},
			Containers: map[string]ct.Container{
				"head":   &fetch.Container{URL: "www.qq.com"},
				"header": &fetch.Container{URL: "www.baidu.com"},
				"menu":   &fetch.Container{URL: "www.baidu.com", Async: true, TTL: 5},
				"content": &fetch.Container{URL: "www.baidu.com", BeforeFetch: func(r *http.Request, child *http.Request) error {
					child.Header.Set("My-Header", "1")

					values := r.URL.Query()
					values.Add("My-Param", "2")
					child.URL.RawQuery = values.Encode()
					return nil
				}},
				"footer": &fetch.Container{URL: "www.163.com"},
			},
			LayoutFunc: templates.Layout,
		},
	}
}
