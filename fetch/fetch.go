package fetch

import (
	"net/http"

	ct "github.com/theplant/containers"
)

type Fetcher struct {
	URL         string
	Primary     bool
	Async       bool
	TTL         int
	BeforeFetch func(r *http.Request, child *http.Request) (err error)
}

func (c *Fetcher) Render(r *http.Request) (html string, err error) {
	return
}

type Results struct {
	values map[string]string
}

func (rs *Results) Get(name string) string {
	if v, ok := rs.values[name]; ok {
		return v
	}
	return ""
}

type Group struct {
	CMap        map[string]ct.Container
	LayoutFunc  func(values *Results) string
	BeforeFetch func(r *http.Request, child *http.Request) (err error)
}

func (c *Group) Render(r *http.Request) (html string, err error) {
	return
}

func (c *Group) Containers(r *http.Request) (cs []ct.Container, err error) {
	return []ct.Container{
		c,
	}, nil
}
