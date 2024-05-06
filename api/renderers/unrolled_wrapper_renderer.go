package renderers

import (
	"html/template"
	"io"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/unrolled/render"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	layoutFileName = "layout"
)

type UnrolledWrapperRenderer struct {
	rnd *render.Render
}

func (urwr *UnrolledWrapperRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return urwr.rnd.HTML(w, 0, name, data)
}

func NewUnrolledWrapperRenderer(directory string, debug bool) *UnrolledWrapperRenderer {
	tp := &UnrolledWrapperRenderer{
		rnd: render.New(render.Options{
			Layout:          layoutFileName,
			Directory:       directory,
			IsDevelopment:   debug,
			RequirePartials: true,
			Funcs: []template.FuncMap{
				{
					"Now": func() int {
						return time.Now().Year()
					},
				},
				{
					"Capitalize": func(str string) string {
						caser := cases.Title(language.English)
						return caser.String(str)
					},
				},
			},
		}),
	}
	return tp
}
