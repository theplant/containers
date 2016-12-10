package reloading

import (
	"fmt"
	"net/http"

	ct "github.com/theplant/containers"
	cb "github.com/theplant/containers/combinators"
)

type reloadEvent struct {
	c         ct.Container
	eventName string
}

func (re *reloadEvent) EventName() string {
	return re.eventName
}

func (re *reloadEvent) Render(r *http.Request) (html string, err error) {
	chtml, err := re.c.Render(r)
	if err != nil {
		return
	}
	html = fmt.Sprintf("<div data-container-reloadon=\"%s\">%s</div>", re.eventName, chtml)
	return
}

func WithReloadEvent(event string, container ct.Container) ct.Container {
	return &reloadEvent{container, event}
}

func OnlyOnReload(container ct.Container) ct.Container {
	return cb.ToContainer(func(r *http.Request) (html string, err error) {
		h := r.Header.Get("Accept")
		if h != "application/x-container-list" {
			return "waiting for reload...", nil
		}
		return container.Render(r)
	})
}
