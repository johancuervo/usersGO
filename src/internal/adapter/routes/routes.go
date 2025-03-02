package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/johancuervo/usersGO/internal/adapter/http"
)

func RegisterRoutes(app *fiber.App, userHandler *http.UserHandler) {
	userGroup := app.Group("/users") // Agrupar rutas bajo `/users`
	userGroup.Get("/", userHandler.GetAllUsers)
	userGroup.Get("/:userId", userHandler.FindUserById)
	userGroup.Post("/", userHandler.CreateUser)
	userGroup.Put("/:userId", userHandler.UpdateUser)
	userGroup.Delete("/:userId", userHandler.DeleteUser)

}
