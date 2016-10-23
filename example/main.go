package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/theplant/appkit/db"
	"github.com/theplant/appkit/log"
	"github.com/theplant/containers"
	"github.com/theplant/containers/example/pages"
	"github.com/theplant/containers/example/parts"
)

func main() {
	dbCfg := db.Config{
		Dialect: "postgres",
		Params:  "sslmode=disable",
	}

	logger := log.Default()
	db, err := db.New(logger, dbCfg)

	if err != nil {
		panic(err)
	}

	c := &container{}
	db.AutoMigrate(c)

	Admin := admin.New(&qor.Config{DB: db})
	Admin.SetSiteName("Containers")

	Admin.AddResource(c)

	mux := http.NewServeMux()

	mux.Handle("/products", containers.Handler(&pages.ProductPage{}, parts.MainLayout))
	mux.Handle("/", containers.Handler(&pages.HomePage{}, parts.MainLayout))
	mux.Handle("/qor", containers.Handler(&QorPage{db}, parts.MainLayout))

	Admin.MountTo("/admin", mux)

	logger.Crit().Log(
		"err", (http.ListenAndServe(":9000", mux)),
	)
}

type container struct {
	gorm.Model
	Element string
	Text    string
}

func (c *container) Render(r *http.Request) (html string, err error) {
	return fmt.Sprintf("<%s>%s<%s>", c.Element, c.Text, c.Element), nil
}

type QorPage struct {
	db *gorm.DB
}

func (p *QorPage) Containers(req *http.Request) (cs []containers.Container, err error) {
	c := []*container{}
	err = p.db.Find(&c).Error
	cs = make([]containers.Container, len(c))
	for i, c1 := range c {
		cs[i] = c1
	}
	return
}
