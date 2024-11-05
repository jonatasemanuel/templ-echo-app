package main

import (
	"net/http"

	"github.com/jonatasemanuel/templ-echo-app/internal/view/page"
	"github.com/labstack/echo/v4"
)

func (app *application) home(c echo.Context) error {
	return render(c, http.StatusOK, page.Home("caju"))
}

func (app *application) hello(c echo.Context) error {
	name := c.Param("name")
	// c.Logger().Info(name) // debuger
	return render(c, http.StatusOK, page.Hello(name))
}
