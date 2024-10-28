package template

import (
	"html/template" // Changed from text/template to html/template for security
	"io"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

// Template helper functions
var funcMap = template.FuncMap{
	// Date formatting
	"formatDate": func(date time.Time) string {
		return date.Format("January 2, 2006")
	},
	// Simple math for pagination
	"add": func(a, b int) int {
		return a + b
	},
	"subtract": func(a, b int) int {
		return a - b
	},
	// String operations
	"lower": strings.ToLower,
	"upper": strings.ToUpper,
	// Slice helpers
	"first": func(arr []interface{}) interface{} {
		if len(arr) > 0 {
			return arr[0]
		}
		return nil
	},
	// Safe HTML content (for blog posts)
	"safe": func(html string) template.HTML {
		return template.HTML(html)
	},
}

func NewTemplateRenderer(e *echo.Echo, paths ...string) {
	tmp1 := template.New("").Funcs(funcMap)
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
