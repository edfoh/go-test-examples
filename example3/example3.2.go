package example3

import (
	"bytes"
	"text/template"
)

const (
	markdownTemplate = `# {{ .Title }}

{{ .Body }}

**References**
{{range $element := .References}}
* {{$element}}{{end}}`
)

type Markdown struct {
	Title      string
	Body       string
	References []string
}

func Generate(md *Markdown) (string, error) {
	b := new(bytes.Buffer)

	t := template.Must(template.New("").Parse(markdownTemplate))
	err := t.Execute(b, md)

	return b.String(), err
}
