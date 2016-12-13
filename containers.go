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

/*
ContainerFunc is a short cut to build simple container that don't depend outside inputs.
Use `combinators.ToContainer` to convert a `ContainerFunc` to a `Container`.
*/
type ContainerFunc func(r *http.Request) (html string, err error)

/*
Page is anything that can return a list of `Containers`, use `containers.PageHandler` to convert a `Page` to a `http.Handler`, So that you can mount it to a url.
*/
type Page interface {
	Containers(r *http.Request) (cs []Container, err error)
}

/*
PageFunc is a short cut to build simple page that don't depend outside inputs.
Use `combinators.ToPage` to convert a `PageFunc` to a `Page`
*/
type PageFunc func(r *http.Request) (cs []Container, err error)

/*
Layout is like a `Container`, but takes another parameter body. use `containers.PageHandler` to combine a `Page` and a `Layout`, and mount to a url.
*/
type Layout func(r *http.Request, body string) (html string, err error)
