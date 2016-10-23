package tests

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/theplant/containers/example/pages"
)

func tsServer() *httptest.Server {
	mux := http.NewServeMux()
	pages.AddRoutes(mux)
	return httptest.NewServer(mux)
}

func bodyString(res *http.Response) (r string) {
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	r = string(b)
	return
}

func TestHome(t *testing.T) {
	ts := tsServer()
	defer ts.Close()

	res, err := http.Get(ts.URL + "/products")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bodyString(res))
}
