package containers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Attrs map[string]string

func Wrap(c Container, el string, attrs Attrs) Container {
	return ContainerFunc(func(r *http.Request) (string, error) {
		out, err := c.Render(r)
		if err != nil {
			return "", err
		} else {
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
		}
	})
}

func FileContainer(filename string) Container {
	return ContainerFunc(func(r *http.Request) (string, error) {
		b, err := ioutil.ReadFile(filename)
		return string(b), err
	})
}

func ScriptByFile(filename string) Container {
	return Wrap(FileContainer(filename), "script", Attrs{"type": "text/javascript"})
}

func StringContainer(value string) Container {
	return ContainerFunc(func(r *http.Request) (string, error) {
		return value, nil
	})
}

func ScriptByString(text string) Container {
	return Wrap(StringContainer(text), "script", Attrs{"type": "text/javascript"})
}
