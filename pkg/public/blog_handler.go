package public

import (
	"net/http"
	"strconv"

	"github.com/AMANSRI99/aman-blogs/pkg/models"
	"github.com/labstack/echo/v4"
)

const (
	defaultPageSize = 10
	defaultOrder    = "created_at"
)

type BlogHandler struct {
	postRepo models.PostRepository
}

func NewBlogHandler(postRepo models.PostRepository) *BlogHandler {
	return &BlogHandler{
		postRepo: postRepo,
	}
}

func (h *BlogHandler) Register(g *echo.Group) {
	g.GET("/posts", h.GetPosts)
	g.GET("/posts/:slug", h.GetPost)
	g.GET("/tags", h.GetTags)
	g.GET("/posts/tag/:tag", h.GetPostsByTag)
}

func (h *BlogHandler) GetPosts(c echo.Context) error {
	ctx := c.Request().Context()

	// Parse page number
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	// Get posts
	posts, err := h.postRepo.GetByStatus(
		ctx,
		models.Published,
		page,
		defaultPageSize,
	)
	if err != nil {
		if c.Request().Header.Get("HX-Request") == "true" {
			return c.HTML(http.StatusInternalServerError, "<p>Error loading posts</p>")
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch posts",
		})
	}

	// Get total count for pagination
	status := models.Published
	total, err := h.postRepo.GetPostCount(ctx, &status, nil)
	if err != nil {
		return err
	}

	// If it's an HTMX request, render the post_list template
	if c.Request().Header.Get("HX-Request") == "true" {
		return c.Render(http.StatusOK, "post_list", map[string]interface{}{
			"Posts":    posts,
			"Page":     page,
			"PageSize": defaultPageSize,
			"Total":    total,
		})
	}

	// Otherwise return JSON (for API calls)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"posts":    posts,
		"page":     page,
		"pageSize": defaultPageSize,
		"total":    total,
	})
}

func (h *BlogHandler) GetPost(c echo.Context) error {
	ctx := c.Request().Context()
	slug := c.Param("slug")

	post, err := h.postRepo.GetBySlug(ctx, slug)
	if err != nil {
		return c.Render(http.StatusNotFound, "error", map[string]interface{}{
			"Error": "Post not found",
		})
	}

	// Get related posts
	related, err := h.postRepo.GetRelatedPosts(ctx, post.ID, 3)
	if err != nil {
		related = []models.Post{} // Empty slice if error
	}

	return c.Render(http.StatusOK, "post", map[string]interface{}{
		"Post":    post,
		"Related": related,
	})
}

func (h *BlogHandler) GetTags(c echo.Context) error {
	ctx := c.Request().Context()

	// Get all published posts to extract tags
	posts, err := h.postRepo.GetByStatus(ctx, models.Published, 1, 100)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch tags",
		})
	}

	// Collect unique tags
	tagMap := make(map[string]int) // tag -> count
	for _, post := range posts {
		for _, tag := range post.Tags {
			tagMap[tag]++
		}
	}

	// Convert map to slice of structs
	type TagCount struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	}
	tags := make([]TagCount, 0, len(tagMap))
	for tag, count := range tagMap {
		tags = append(tags, TagCount{Name: tag, Count: count})
	}

	return c.JSON(http.StatusOK, tags)
}

func (h *BlogHandler) GetPostsByTag(c echo.Context) error {
	ctx := c.Request().Context()
	tag := c.Param("tag")

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	posts, err := h.postRepo.GetByTag(
		ctx,
		tag,
		page,
		defaultPageSize,
	)
	if err != nil {
		if c.Request().Header.Get("HX-Request") == "true" {
			return c.HTML(http.StatusInternalServerError, "<p>Error loading posts</p>")
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch posts",
		})
	}

	if c.Request().Header.Get("HX-Request") == "true" {
		return c.Render(http.StatusOK, "post_list", map[string]interface{}{
			"Posts":    posts,
			"Page":     page,
			"PageSize": defaultPageSize,
			"Tag":      tag,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"posts":    posts,
		"page":     page,
		"pageSize": defaultPageSize,
		"tag":      tag,
	})
}
