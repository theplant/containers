package reloading

import (
	"fmt"
	"net/http"

	"github.com/theplant/containers"
)

func WithReloadEvent(event string, container containers.Container) containers.Container {
	return containers.ContainerFunc(func(r *http.Request) (html string, err error) {
		c, err := container.Render(r)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("<div data-container-reloadon=\"%s\">%s</div>", event, c), nil
	})
}

func OnlyOnReload(container containers.Container) containers.Container {
	return containers.ContainerFunc(func(r *http.Request) (html string, err error) {
		h := r.Header.Get("Accept")
		if h != "application/x-container-list" {
			return "waiting for reload...", nil
		}
		return container.Render(r)
	})
}
