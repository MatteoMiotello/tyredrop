package email

import (
	"bytes"
	"pillowww/titw/internal/email/templates"
	"text/template"
)

type Templater struct {
	TemplateBase string
}

func NewTemplater() *Templater {
	return &Templater{
		TemplateBase: "./internal/email/templates/base.tmpl",
	}
}

func (r Templater) Process(eTemplate string, params *templates.EmailParams) (*bytes.Buffer, error) {
	tFile := "./internal/email/templates/" + eTemplate + ".tmpl"

	b := new(bytes.Buffer)

	files, err := template.ParseFiles(r.TemplateBase, tFile)

	if err != nil {
		return nil, err
	}

	err = files.Execute(b, *params)

	if err != nil {
		return nil, err
	}

	return b, nil
}
