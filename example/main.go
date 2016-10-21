package main

import (
	"log"
	"net/http"

	"github.com/theplant/containers"
)

func main() {
	http.Handle("/", &containers.MainHandler{})
	log.Fatal(http.ListenAndServe(":9000", nil))
}
