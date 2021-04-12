package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nid90/kawaii-blog-engine/handlers"
)

func SetupRoutes(fiberApp *fiber.App) {
	// middleware
	app := fiberApp.Group("/", logger.New())

	// routes
	postGroup := app.Group("/posts")
	postGroup.Get("/", handlers.FetchPosts)
	postGroup.Post("/", handlers.CreatePost)
	postGroup.Get("/new", handlers.NewPost)
}
