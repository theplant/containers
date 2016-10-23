package pages

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/theplant/containers"
	r "github.com/theplant/containers/reloading"
)

type HomePage struct {
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

func text(text string) containers.Container {
	return toC(func(r *http.Request) (string, error) {
		return text, nil
	})
}

func (hp *HomePage) Containers(req *http.Request) (cs []containers.Container, err error) {
	return []containers.Container{
		text("<script src=\"https://cdnjs.cloudflare.com/ajax/libs/fetch/1.0.0/fetch.min.js\"></script>"),
		text("<a href=\"/products\">products</a>"),
		text("triggers `a`, no reload"),
		makeContainer(rand.Int(), "a"),
		text("triggers `b`, no reload"),
		makeContainer(rand.Int(), "b"),

		text("<h1>reload on `a`</h1>"),
		text("triggers `b`"),
		r.Reloadable("a", makeContainer(rand.Int(), "b")),
		text("triggers `a`"),
		r.Reloadable("a", makeContainer(rand.Int(), "a")),
		r.Reloadable("a", text(fmt.Sprintf("static: %d", rand.Int()))),

		text("<h1>reload on `b`</h1>"),
		text("triggers `b`"),
		r.Reloadable("b", makeContainer(rand.Int(), "b")),
		text("triggers `a`"),
		r.Reloadable("b", makeContainer(rand.Int(), "a")),
		r.Reloadable("b", text(fmt.Sprintf("static: %d", rand.Int()))),

		r.Reloadable("b", r.OnlyOnReload(text("reloaded!"))),

		r.ReloadingScript(),
		containers.Script("script.js"),
	}, nil
}
