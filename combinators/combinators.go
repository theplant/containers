package containers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	ct "github.com/theplant/containers"
)

func ContainerFunc(f func(r *http.Request) (html string, err error)) ct.Container {
	return containerFunc{f}
}

type containerFunc struct {
	cf func(r *http.Request) (html string, err error)
}

func (f containerFunc) Render(r *http.Request) (html string, err error) {
	return f.cf(r)
}

type pageFunc struct {
	pf func(r *http.Request) (cs []ct.Container, err error)
}

func (p pageFunc) Containers(r *http.Request) (cs []ct.Container, err error) {
	return p.pf(r)
}

func PageFunc(f func(r *http.Request) (cs []ct.Container, err error)) ct.Page {
	return pageFunc{f}
}

type Attrs map[string]string

func Wrap(c ct.Container, el string, attrs Attrs) ct.Container {
	return ContainerFunc(func(r *http.Request) (string, error) {
		out, err := c.Render(r)
		if err != nil {
			return "", err
		}

		attrsbuf := bytes.NewBuffer(nil)
		if attrs != nil {
			for key, value := range attrs {
				attrsbuf.WriteString(" ")
				attrsbuf.WriteString(key)
				attrsbuf.WriteString(`="`)
				attrsbuf.WriteString(value)
				attrsbuf.WriteString(`"`)
			}
		}
		return fmt.Sprintf("<%s%s>%s</%s>\n", el, attrsbuf.String(), out, el), nil
	})
}

func FileContainer(filename string) ct.Container {
	return ContainerFunc(func(r *http.Request) (string, error) {
		b, err := ioutil.ReadFile(filename)
		return string(b), err
	})
}

func ScriptByFile(filename string) ct.Container {
	return Wrap(FileContainer(filename), "script", Attrs{"type": "text/javascript"})
}

func StringContainer(value string) ct.Container {
	return ContainerFunc(func(r *http.Request) (string, error) {
		return value, nil
	})
}

func ScriptByString(text string) ct.Container {
	return Wrap(StringContainer(text), "script", Attrs{"type": "text/javascript"})
}
