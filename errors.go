package containers

import (
	"context"
	"fmt"
	"net/http"
)

type ErrHandler interface {
	HandleErr(w http.ResponseWriter, r *http.Request, err error)
}

type redirectErrorHandler struct {
	url  string
	code int
}

func (reh *redirectErrorHandler) HandleErr(w http.ResponseWriter, r *http.Request, err error) {
	http.Redirect(w, r, reh.url, reh.code)
}

func (reh *redirectErrorHandler) Error() string {
	return fmt.Sprintf("%d: redirect to: %s", reh.code, reh.url)
}

func NewRedirectError(url string, code int) (err error) {
	err = &redirectErrorHandler{url: url, code: code}
	return
}

const internalServerErrorHandlerKey = "containers.InternalServerErrorHandler"

type internalServerErrorHandler struct {
	h  http.Handler
	eh ErrHandler
}

func (iseh *internalServerErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r = r.WithContext(context.WithValue(r.Context(), internalServerErrorHandlerKey, iseh.eh))
	iseh.h.ServeHTTP(w, r)
}

func UseErrHandler(h http.Handler, errhandler ErrHandler) (handler http.Handler) {
	handler = &internalServerErrorHandler{h: h, eh: errhandler}
	return
}
