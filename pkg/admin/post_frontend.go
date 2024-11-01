package admin

import (
	"net/http"
	"strings"

	"github.com/AMANSRI99/aman-blogs/pkg/models"
	"github.com/labstack/echo/v4"
)

type PostFrontend struct {
	repo models.PostRepository
}

func NewPostFrontend(g *echo.Group, repo models.PostRepository) {
	fe := PostFrontend{
		repo: repo,
	}
	g.GET("/posts/new", fe.New)
	g.POST("/posts", fe.Create)
	g.GET("/posts/new/cancel", fe.Cancel)
}

func (fe *PostFrontend) New(c echo.Context) error {
	return c.Render(http.StatusOK, "posts_new", nil)
}

func (fe *PostFrontend) Create(c echo.Context) error {
	ctx := c.Request().Context()

	// Parse form data
	title := strings.TrimSpace(c.FormValue("title"))
	content := strings.TrimSpace(c.FormValue("content"))
	status := models.PostStatus(c.FormValue("status"))
	tagsStr := c.FormValue("tags")

	// Validation
	if title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Title is required",
		})
	}

	if content == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Content is required",
		})
	}

	// If status is empty, default to draft
	if status == "" {
		status = models.Draft
	}

	// Process tags
	tags := []string{}
	if tagsStr != "" {
		for _, tag := range strings.Split(tagsStr, ",") {
			trimmedTag := strings.TrimSpace(tag)
			if trimmedTag != "" {
				tags = append(tags, trimmedTag)
			}
		}
	}

	// Create post
	post := &models.Post{
		Title:   title,
		Content: content,
		Status:  status,
		Tags:    tags,
	}

	// Save post
	err := fe.repo.Upsert(ctx, post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create post",
		})
	}

	// Check if it's an HTMX request
	if c.Request().Header.Get("HX-Request") == "true" {
		// Return a success message for HTMX
		return c.HTML(http.StatusOK, `
			<div class="bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded relative" role="alert">
				<strong class="font-bold">Success!</strong>
				<span class="block sm:inline"> Post created successfully.</span>
			</div>
		`)
	}

	// Regular form submission - redirect to admin index
	return c.Redirect(http.StatusSeeOther, "/admin")
}

func (fe *PostFrontend) Cancel(c echo.Context) error {
	// If it's an HTMX request, return to admin content
	if c.Request().Header.Get("HX-Request") == "true" {
		posts, err := fe.repo.Get(c.Request().Context(), 1, 10, "created_at", "DESC")
		if err != nil {
			return err
		}
		return c.Render(http.StatusOK, "admin_index", map[string]interface{}{
			"Posts": posts,
		})
	}

	// Regular request - redirect to admin index
	return c.Redirect(http.StatusSeeOther, "/admin")
}
