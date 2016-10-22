package pages

import (
	"fmt"
	"net/http"

	"github.com/theplant/containers"
)

type HomePage struct {
}

func makeContainer(label int) containers.Container {
	return func(r *http.Request) (html string, err error) {
		return fmt.Sprintf("<div>%d</div>", label), nil
	}
}

func repeat(c containers.Container) containers.Container {
	return func(r *http.Request) (html string, err error) {
		out, _ := c(r)
		out2, _ := c(r)
		return out + out2, nil
	}
}

func (hp *HomePage) Containers(r *http.Request) (cs []containers.Container, err error) {
	return []containers.Container{
		repeat(makeContainer(1)),
		makeContainer(2),
		makeContainer(3),
	}, nil
}
