// pkg/public/blog_handler.go
package public

import (
	"net/http"
	"strconv"

	"github.com/AMANSRI99/aman-blogs/pkg/models"
	"github.com/AMANSRI99/aman-blogs/utils/scopes"
	"github.com/labstack/echo/v4"
)

type BlogHandler struct {
	postRepo models.PostRepository
}

func NewBlogHandler(postRepo models.PostRepository) *BlogHandler {
	return &BlogHandler{
		postRepo: postRepo,
	}
}

// Register registers all blog routes
func (h *BlogHandler) Register(g *echo.Group) {
	g.GET("/posts", h.GetPosts)
	g.GET("/posts/:slug", h.GetPost)
	g.GET("/tags", h.GetTags)
	g.GET("/posts/tag/:tag", h.GetPostsByTag)
}

// pkg/public/blog_handler.go
func (h *BlogHandler) GetPosts(c echo.Context) error {
	ctx := c.Request().Context()

	// Convert page parameter from string to int
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1 // Default to first page if invalid or not provided
	}

	posts, err := h.postRepo.Get(
		ctx,
		page,                      // page
		10,                        // pageSize
		"created_at",              // orderBy
		string(scopes.Descending), // orderDir
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch posts",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"posts": posts,
		"page":  page,
	})
}
func (h *BlogHandler) GetPost(c echo.Context) error {
	ctx := c.Request().Context()

	// Convert id from string to uint
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid post ID",
		})
	}
	id := uint(id64)

	post, err := h.postRepo.GetById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Post not found",
		})
	}

	return c.Render(http.StatusOK, "post", map[string]interface{}{
		"post": post,
	})
}

func (h *BlogHandler) GetTags(c echo.Context) error {
	// Implementation pending tag functionality
	return c.JSON(http.StatusOK, []string{})
}

func (h *BlogHandler) GetPostsByTag(c echo.Context) error {
	// Implementation pending tag functionality
	return c.JSON(http.StatusOK, []string{})
}
