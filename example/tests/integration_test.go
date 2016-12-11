package tests

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"encoding/json"

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
		t.Error(err)
	}
	body := bodyString(res)
	if strings.Index(body, "addToCart") < 0 {
		t.Error(body)
	}
}

func TestHome(t *testing.T) {
	ts := httptest.NewServer(reloading.ReloadablePageHandler(&pages.HomePage{}, parts.MainLayout))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	body := bodyString(res)
	if strings.Index(body, "data-container-id") < 0 {
		t.Error(body)
	}
}

func TestReloadNestedContainers(t *testing.T) {
	ts := httptest.NewServer(reloading.ReloadablePageHandler(&pages.ProductPage{}, parts.MainLayout))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	body := bodyString(res)
	if strings.Index(body, "2.1.1") < 0 {
		t.Error(body)
	}
	req, err := http.NewRequest("GET", ts.URL+"?containersByTags=cart_updated", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Accept", "application/x-container-list")
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	body = bodyString(res)
	var values map[string]string
	json.Unmarshal([]byte(body), &values)
	if values["2.1.1"] == "" || values["2.2"] == "" {
		t.Error(body)
	}
}

func TestStructuredPage(t *testing.T) {
	ts := httptest.NewServer(containers.PageHandler(&pages.StructuredPage{}, parts.MainLayout))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	body := bodyString(res)
	// t.Error(body)
	if strings.Index(body, "sidemenu") < 0 {
		t.Error(body)
	}
}
