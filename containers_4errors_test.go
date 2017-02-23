package containers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"fmt"

	ct "github.com/theplant/containers"
	cb "github.com/theplant/containers/combinators"
	"github.com/theplant/testingutils"
)

type errhandler struct {
}

func (eh *errhandler) HandleErr(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "err `%s` is properly handled", err.Error())
}

func Cat(r *http.Request) (html string, err error) {
	err = ct.NewRedirectError("/fff", http.StatusPermanentRedirect)
	return
}

func BadBoy(r *http.Request) (html string, err error) {
	err = errors.New("ohh No.")
	return
}

func MyCatHome(r *http.Request) (cs []ct.Container, err error) {
	cs = []ct.Container{
		cb.ToContainer(Cat),
	}
	return
}

func ExampleContainer_4errors() {

	http.Handle("/page4", ct.UseErrHandler(ct.PageHandler(cb.ToPage(MyCatHome), nil), &errhandler{}))
	//Output:

}

func TestErrorRedirect(t *testing.T) {
	ts := httptest.NewServer(ct.UseErrHandler(ct.PageHandler(cb.ToPage(MyCatHome), nil), &errhandler{}))
	defer ts.Close()

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	res, err := client.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusPermanentRedirect {
		t.Error("wrong http status", res.Status)
	}
}

func TestErrorInternal(t *testing.T) {
	ts := httptest.NewServer(ct.UseErrHandler(ct.PageHandler(cb.ToPage(func(r *http.Request) (cs []ct.Container, err error) {
		cs = []ct.Container{
			cb.ToContainer(BadBoy),
		}
		return
	}), nil), &errhandler{}))

	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusInternalServerError {
		t.Error("didn't handle err")
	}
	expectedBody := "err `ohh No.` is properly handled"
	diff := testingutils.PrettyJsonDiff(expectedBody, res.Body)
	if len(diff) > 0 {
		t.Error(diff)
	}

}
