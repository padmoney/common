package simplerender

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

type Template struct {
	path    string
	content string
	loaded  bool
}

func NewTemplate(path string) Template {
	return Template{path: path}
}

func renderVariable(content, variable, value string) string {
	pattern := fmt.Sprintf("{{ *%s *}}", variable)
	compiled := regexp.MustCompile(pattern)
	return compiled.ReplaceAllString(content, value)
}

func (t *Template) read() error {
	if t.loaded {
		return nil
	}

	content, err := ioutil.ReadFile(t.path)
	if err != nil {
		return err
	}
	t.content = string(content)
	t.loaded = true
	return nil
}

func (t *Template) Render(params map[string]string) (string, error) {
	err := t.read()
	if err != nil {
		return "", err
	}

	content := t.content
	for k, v := range params {
		content = renderVariable(content, k, v)
	}
	return content, nil
}
