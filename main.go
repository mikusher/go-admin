package main

import (
	"github.com/gofiber/fiber/v2"
	"go-admin/database"
	"go-admin/routes"
)

func main() {
	// connect to database
	database.Connect()

	// setup app
	app := fiber.New()
	routes.Setup(app)

	//listener port
	app.Listen(":8000")
}
