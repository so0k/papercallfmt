package main

import (
	"io"
	"path"
	"text/template"

	"github.com/Masterminds/sprig"
)

// RenderTemplate renders Team using tplFile into target Writer
func RenderTemplate(s *Submission, tplFile string, wr io.Writer) error {
	tpl, err := template.New(path.Base(tplFile)).
		Funcs(sprig.TxtFuncMap()).
		ParseFiles(tplFile)
	if err != nil {
		return err
	}
	return tpl.Execute(wr, s)
}
