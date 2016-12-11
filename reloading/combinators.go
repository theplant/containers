package reloading

import (
	"net/http"

	ct "github.com/theplant/containers"
	cb "github.com/theplant/containers/combinators"
)

type tagc struct {
	c        ct.Container
	tagNames string
}

func (re *tagc) TagNames() string {
	return re.tagNames
}

func (re *tagc) Render(r *http.Request) (html string, err error) {
	return re.c.Render(r)
}

func WithTags(tagNames string, container ct.Container) ct.Container {
	return &tagc{container, tagNames}
}

func OnlyOnReload(container ct.Container) ct.Container {
	return cb.ToContainer(func(r *http.Request) (html string, err error) {
		h := r.Header.Get("Accept")
		if h != "application/x-container-list" {
			return
		}
		return container.Render(r)
	})
}
