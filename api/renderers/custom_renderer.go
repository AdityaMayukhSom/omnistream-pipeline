package renderers

import (
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
)

const (
	templateFileExtention = ".go.html"
)

// A structure implementing the echo.Rendered interface for templates rendering.
//
// Deprecated: Use UnrolledWrappedRenderer instead.
// See: https://github.com/unrolled/render?tab=readme-ov-file#echo
type CustomRenderer struct {
	template *template.Template
	debug    bool
	location string
}

func (t *CustomRenderer) ReloadTemplates() {
	t.template = template.New("").Funcs(template.FuncMap{
		"Now": func() int {
			return time.Now().Year()
		},
	})

	cleanRoot := filepath.Clean(t.location)

	walkFunc := func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(path, templateFileExtention) {
			if _, err = t.template.ParseFiles(path); err != nil {
				log.Error(err)
			}
		}
		return err
	}

	if err := filepath.Walk(cleanRoot, walkFunc); err != nil {
		log.Error(err)
	}
}

func (t *CustomRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if t.debug {
		t.ReloadTemplates()
	}

	return t.template.ExecuteTemplate(w, name, data)
}

// Deprecated: Use UnrolledWrappedRenderer instead.
// See: https://github.com/unrolled/render?tab=readme-ov-file#echo
func NewCustomRenderer(directory string, debug bool) *CustomRenderer {
	tp := &CustomRenderer{
		location: directory,
		debug:    debug,
	}
	tp.ReloadTemplates()
	return tp
}
