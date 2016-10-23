package containers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Wrap(c Container, el string) Container {
	return ContainerFunc(func(r *http.Request) (string, error) {
		out, err := c.Content(r)
		if err != nil {
			return "", err
		} else {
			return fmt.Sprintf("<%s>%s</%s>", el, out, el), nil
		}
	})
}

func FileContainer(filename string) Container {
	return ContainerFunc(func(r *http.Request) (string, error) {
		b, err := ioutil.ReadFile(filename)
		return string(b), err
	})
}

func Script(filename string) Container {
	return Wrap(FileContainer(filename), "script")

}
