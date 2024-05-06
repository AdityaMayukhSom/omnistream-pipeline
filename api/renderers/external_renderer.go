package renderers

import (
	"github.com/labstack/echo/v4"
)

// Returns implementation of the echo renderer interface
func NewRenderer(directory string, debug bool) echo.Renderer {
	return NewUnrolledWrapperRenderer(directory, debug)
}
