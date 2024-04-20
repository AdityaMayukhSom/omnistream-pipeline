package api

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Renderer struct {
	template *template.Template
	debug    bool
	location string
}

func (t *Renderer) ReloadTemplates() {
	t.template = template.Must(template.ParseGlob(t.location))
}

func (t *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if t.debug {
		t.ReloadTemplates()
	}

	return t.template.ExecuteTemplate(w, name, data)
}

func NewRenderer(location string, debug bool) *Renderer {
	tp := &Renderer{
		location: location,
		debug:    debug,
	}
	tp.ReloadTemplates()
	return tp
}
