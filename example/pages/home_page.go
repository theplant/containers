package pages

import (
	"context"
	"fmt"
	"net/http"

	"github.com/theplant/containers"
)

type HomePage struct {
}

func makeContainer(label int) containers.Container {
	return func(r *http.Request, ctx context.Context) (html string, err error) {
		return fmt.Sprintf("<div>%d</div>", label), nil
	}
}

func (hp *HomePage) Containers(r *http.Request, ctx context.Context) (cs []containers.Container, err error) {

	return []containers.Container{
		makeContainer(1),
		makeContainer(2),
		makeContainer(3),
	}, nil
}
