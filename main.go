package main

import (
	"fmt"
	"net/http"

	"github.com/AMANSRI99/aman-blogs/config"
	"github.com/AMANSRI99/aman-blogs/template"
	"github.com/AMANSRI99/aman-blogs/utils/routing"
	"github.com/labstack/echo/v4"
)

func main() {
	println("welcome to my nook of the internet")

	e := echo.New()
	routing.SetupRouter(e)
	e.Static("/dist", "dist")

	template.NewTemplateRenderer(e, []string{
		"public/*.html",
	})

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", "Sample template")
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Get(config.PORT))))
}
