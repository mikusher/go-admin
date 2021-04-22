package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-admin/controllers"
	"go-admin/middlewares"
)

func Setup(app *fiber.App) {
	//public router
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	//activate middleware
	app.Use(middlewares.IsAuthenticated)

	//private router need middleware
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
}
