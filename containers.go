/*
Package containers provide better structures for building web applications
*/
package containers

import "net/http"

/*
Container is a build block for html pape. it's an isolated block that should contains how to prepare data, and how to render itself to html.

By following container interface to implements parts of your html page. You can utilize package built upon containers interface that implements specific features, like reloading, combinators package.
*/
type Container interface {
	Render(r *http.Request) (html string, err error)
}

type ContainerFunc func(r *http.Request) (html string, err error)

type Page interface {
	Containers(r *http.Request) (cs []Container, err error)
}

type PageFunc func(r *http.Request) (cs []Container, err error)

type Layout func(r *http.Request, body string) (html string, err error)
