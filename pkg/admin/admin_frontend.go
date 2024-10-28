package admin

import (
	"github.com/AMANSRI99/aman-blogs/pkg/models"
	"github.com/AMANSRI99/aman-blogs/utils/scopes"
	"github.com/labstack/echo/v4"
)

type AdminFrontend struct {
	repo models.PostRepository
}

func NewAdminFrontend(g *echo.Group, repo models.PostRepository) {
	fe := AdminFrontend{
		repo: repo,
	}
	g.GET("", fe.Index)
}

func (fe *AdminFrontend) Index(c echo.Context) error {
	ctx := c.Request().Context()

	posts, err := fe.repo.Get(
		ctx,
		1,                         // page
		10,                        // pageSize
		"created_at",              // orderBy
		string(scopes.Descending), // orderDir
	)

	if err != nil {
		return err
	}
	resp := map[string]interface{}{
		"Posts": posts,
	}
	return c.Render(200, "admin_index", resp)
}
