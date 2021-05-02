package routes

import (
	"go-admin/controllers"
	"go-admin/middlewares"

	"github.com/gofiber/fiber/v2"
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

	//CRUD
	// getAll users
	app.Get("/api/users", controllers.AllUsers)
	//create user
	app.Post("/api/users", controllers.CreateUser)
	//get user by Id
	app.Get("/api/users/:id", controllers.GetUser)
	//update user by Id
	app.Put("/api/users/:id", controllers.UpdateUser)
	//delete user by Id
	app.Delete("/api/users/:id", controllers.DeleteUser)

	// getAll users
	app.Get("/api/roles", controllers.AllRoles)
	//create user
	app.Post("/api/roles", controllers.CreateRole)
	//get user by Id
	app.Get("/api/roles/:id", controllers.GetRole)
	//update user by Id
	app.Put("/api/roles/:id", controllers.UpdateRole)
	//delete user by Id
	app.Delete("/api/roles/:id", controllers.DeleteRole)

	//permission
	app.Get("/api/permissions", controllers.AllPermission)
}
