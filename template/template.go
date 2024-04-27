package template

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func NewTemplateRenderer(e *echo.Echo, paths []string) {
	tmp1 := &template.Template{}
	for i := range paths {
		template.Must(tmp1.ParseGlob(paths[i]))
	}
	t := newTemplate(tmp1)
	e.Renderer = t
}

func newTemplate(templates *template.Template) echo.Renderer {
	return &Template{
		Templates: templates,
	}
}
