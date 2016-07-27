package demo

import (
	"github.com/labstack/echo"
)

func New() (e *echo.Echo) {
	e = echo.New()
	echoBinder := e.Binder()
	binder := &binder{echoBinder}
	e.SetBinder(binder)
	return
}
