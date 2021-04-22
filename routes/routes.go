package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-admin/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
}
