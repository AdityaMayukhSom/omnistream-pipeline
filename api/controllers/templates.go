package controllers

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

// A structure implementing the echo.Rendered interface for templates rendering.
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

func RenderHelloWorldPage(c echo.Context) error {
	return c.Render(http.StatusOK, "helloworld.go.html", map[string]interface{}{
		"message": "Hello from Rendered Views",
	})
}

func RenderHomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "homepage.go.html", map[string]interface{}{})
}
