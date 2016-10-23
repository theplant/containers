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

func reloadable(event string, container containers.Container) containers.Container {
	return toC(func(r *http.Request) (html string, err error) {
		c, err := container.Content(r)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("<div data-container-reloadon=\"%s\">%s</div>", event, c), nil
	})
}

func onlyOnReload(container containers.Container) containers.Container {
	return toC(func(r *http.Request) (html string, err error) {
		h := r.Header.Get("Accept")
		if h != "application/x-container-list" {
			return "waiting for reload...", nil
		}
		return container.Content(r)
	})
}

func toC(f func(r *http.Request) (html string, err error)) containers.Container {
	return containers.ContainerFunc(f)
}

func makeContainer(label int, event string) containers.Container {
	return toC(func(r *http.Request) (html string, err error) {
		return fmt.Sprintf("<button data-container-event=\"%s\">%d</button>", event, label), nil
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
		text("triggers `a`, no reload"),
		makeContainer(rand.Int(), "a"),
		text("triggers `b`, no reload"),
		makeContainer(rand.Int(), "b"),

		text("<h1>reload on `a`</h1>"),
		text("triggers `b`"),
		reloadable("a", makeContainer(rand.Int(), "b")),
		text("triggers `a`"),
		reloadable("a", makeContainer(rand.Int(), "a")),
		reloadable("a", text(fmt.Sprintf("static: %d", rand.Int()))),

		text("<h1>reload on `b`</h1>"),
		text("triggers `b`"),
		reloadable("b", makeContainer(rand.Int(), "b")),
		text("triggers `a`"),
		reloadable("b", makeContainer(rand.Int(), "a")),
		reloadable("b", text(fmt.Sprintf("static: %d", rand.Int()))),

		reloadable("b", onlyOnReload(text("reloaded!"))),

		wrap(fileContainer("script.js"), "script"),
	}, nil
}
