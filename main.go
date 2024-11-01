package main

import (
	"fmt"
	"net/http"

	"github.com/AMANSRI99/aman-blogs/config"
	"github.com/AMANSRI99/aman-blogs/db"
	"github.com/AMANSRI99/aman-blogs/pkg/admin"
	"github.com/AMANSRI99/aman-blogs/pkg/public"
	"github.com/AMANSRI99/aman-blogs/pkg/repository"
	"github.com/AMANSRI99/aman-blogs/template"
	"github.com/AMANSRI99/aman-blogs/utils/routing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	println("welcome to my nook of the internet")

	e := echo.New()

	// Adding error logging middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routing.SetupRouter(e)
	e.Static("/dist", "dist")

	template.NewTemplateRenderer(e,
		"public/*.html",
		"public/admin/views/*.html",
		"public/admin/components/*.html",
		"public/blog/*.html", // Add blog templates
	)
	// Initializing db connection
	dbConnection := db.GetDatabase()

	// Initializing repository
	postRepo := repository.NewPostRepository(dbConnection)

	// Public routes
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", "Sample template")
	})

	// Blog routes
	blogHandler := public.NewBlogHandler(postRepo) // This creates a handler that will manage blog-related operations
	// It gets postRepo so it can perform database operations. This is called repository pattern.
	blogGroup := e.Group("/blog") // This creates /blog/* routes
	blogHandler.Register(blogGroup)

	// Admin routes
	adminGrp := e.Group("admin")
	admin.NewAdminFrontend(adminGrp, postRepo)
	admin.NewPostFrontend(adminGrp, postRepo)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Get(config.PORT))))
}
