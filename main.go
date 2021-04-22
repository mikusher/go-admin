package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-admin/database"
	"go-admin/routes"
)

func main() {
	// connect to database
	database.Connect()

	// setup app
	app := fiber.New()
	//cors
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	//listener port
	app.Listen(":8000")
}
