package tests

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/theplant/containers"
	"github.com/theplant/containers/example/pages"
	"github.com/theplant/containers/example/parts"
	"github.com/theplant/containers/reloading"
)

func bodyString(res *http.Response) (r string) {
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	r = string(b)
	return
}

func TestPageHandlerCanAcceptNilLayout(t *testing.T) {
	ts := httptest.NewServer(containers.PageHandler(&pages.ProductPage{}, nil))
	defer ts.Close()

	_, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

}

func TestProducts(t *testing.T) {
	ts := httptest.NewServer(reloading.ReloadablePageHandler(&pages.ProductPage{}, parts.MainLayout))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bodyString(res))
}

func TestHome(t *testing.T) {
	ts := httptest.NewServer(reloading.ReloadablePageHandler(&pages.HomePage{}, parts.MainLayout))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bodyString(res))
}
