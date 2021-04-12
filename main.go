package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/nid90/kawaii-blog-engine/database"
	"github.com/nid90/kawaii-blog-engine/models"
	"github.com/nid90/kawaii-blog-engine/routes"
)

func InitMigrations() {
	database.DBConn.AutoMigrate(&models.Post{})
	fmt.Println("Auto-migrated Post")
}

func main() {
	// init
	database.InitDatabase()
	InitMigrations()
	engine := html.New("./views/", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	routes.SetupRoutes(app)

	// start
	app.Listen(":3000")
}
