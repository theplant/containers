package gin

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	ct "github.com/theplant/containers"
)

const ContextKey = "containers.gin.context"

type ContainerFunc func(c *gin.Context) (html string, err error)

type ginContainer struct {
	ginCf ContainerFunc
}

func (c ginContainer) Render(r *http.Request) (html string, err error) {
	ctx := r.Context().Value(ContextKey).(*gin.Context)
	return c.ginCf(ctx)
}

type PageFunc func(c *gin.Context) (cs []ContainerFunc, err error)

type pageFunc struct {
	pf PageFunc
}

func (p pageFunc) Containers(r *http.Request) (cs []ct.Container, err error) {
	ctx := r.Context().Value(ContextKey).(*gin.Context)
	var ginContainers []ContainerFunc
	ginContainers, err = p.pf(ctx)
	if err != nil {
		return
	}
	for _, gc := range ginContainers {
		cs = append(cs, ginContainer{gc})
	}
	return
}

func ToPage(f PageFunc) ct.Page {
	return pageFunc{f}
}

func WrapH(h http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), ContextKey, c))
		h.ServeHTTP(c.Writer, c.Request)
	}
}
