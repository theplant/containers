package gin

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	rl "github.com/theplant/containers/reloading"
)

func ginContainer1(ctx *gin.Context) (html string, err error) {
	html = ctx.Param("name")
	return
}

func ginContainer2(ctx *gin.Context) (html string, err error) {
	html = ctx.PostForm("file")
	return
}

func page(ctx *gin.Context) (cs []ContainerFunc, err error) {
	cs = []ContainerFunc{
		ginContainer1,
		ginContainer2,
	}
	return
}

func TestGinContainers(t *testing.T) {
	engine := gin.Default()
	engine.GET("/products/:name", WrapH(rl.ReloadablePageHandler(ToPage(page), nil)))

	ts := httptest.NewServer(engine)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/products/the_strange_name")
	if err != nil {
		log.Fatal(err)
	}
	body := bodyString(res)
	if strings.Index(body, "the_strange_name") < 0 {
		t.Error(body)
	}
}

func bodyString(res *http.Response) (r string) {
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	r = string(b)
	return
}
