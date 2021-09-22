package users_routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shuklasaharsh/REST-API/pkg/REST/controllers/users"
)

func InitialiseRoutes(app *fiber.App) {
	app.Get("/", )
	app.Post("/users", users_controllers.CreateUser)
	app.Get("/users/all", users_controllers.GetAllUsers)
	app.Get("/users", users_controllers.GetUser)
	app.Put("/users", users_controllers.UpdateUserByID)
	app.Delete("/users", users_controllers.DeleteUser)
}