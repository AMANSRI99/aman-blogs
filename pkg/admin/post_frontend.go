package admin

import (
	"github.com/AMANSRI99/aman-blogs/pkg/models"
	"github.com/labstack/echo/v4"
)

type Postfrontend struct {
	repo models.PostRepository
}

func NewPostfrontend(g *echo.Group, repo models.PostRepository) {
	fe := &Postfrontend{
		repo: repo,
	}

	g.GET("/posts/new", fe.PostsNew)
	g.GET("posts/new/cancel", fe.PostsNewCancel)
}

func (fe *Postfrontend) PostsNew(c echo.Context) error {
	return c.Render(200, "posts_new", nil)
}

func (fe *Postfrontend) PostsNewCancel(c echo.Context) error {
	c.Response().Header().Set("HX-REDIRECT", "/admin")
	return c.Render(200, "admin_index", nil)
}
