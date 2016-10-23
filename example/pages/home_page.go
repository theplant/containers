package pages

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/theplant/containers"
)

type HomePage struct {
}

func toC(f func(r *http.Request) (html string, err error)) containers.Container {
	return containers.ContainerFunc(f)
}

func makeContainer(label int) containers.Container {
	return toC(func(r *http.Request) (html string, err error) {
		return fmt.Sprintf("<button data-container-event=\"a\" data-container-reloadon=\"a\">%d</button>", label), nil
	})
}

func repeat(c containers.Container) containers.Container {
	return toC(func(r *http.Request) (html string, err error) {
		out, _ := c.Content(r)
		out2, _ := c.Content(r)
		return out + out2, nil
	})
}

func wrap(c containers.Container, el string) containers.Container {
	return toC(func(r *http.Request) (string, error) {
		out, err := c.Content(r)
		if err != nil {
			return "", err
		} else {
			return fmt.Sprintf("<%s>%s</%s>", el, out, el), nil
		}
	})
}

func fileContainer(filename string) containers.Container {
	return toC(func(r *http.Request) (string, error) {
		b, err := ioutil.ReadFile(filename)
		return string(b), err
	})
}

func text(text string) containers.Container {
	return toC(func(r *http.Request) (string, error) {
		return text, nil
	})
}

func (hp *HomePage) Containers(r *http.Request) (cs []containers.Container, err error) {
	return []containers.Container{
		text("<script src=\"https://cdnjs.cloudflare.com/ajax/libs/fetch/1.0.0/fetch.min.js\"></script>"),
		text("<a href=\"/products\">products</a>"),
		repeat(makeContainer(1)),
		makeContainer(rand.Int()),
		text(fmt.Sprintf("a random number that doesn't change: %d", rand.Int())),
		makeContainer(rand.Int()),
		text("the next button should trigger events, but not reload itself"),
		text(fmt.Sprintf("<button data-container-event=\"a\">%d</button>", rand.Int())),
		wrap(fileContainer("script.js"), "script"),
	}, nil
}
