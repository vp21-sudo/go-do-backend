package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vp21-sudo/go-do-backend/handlers"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/todos", handlers.GetTodos)
	api.Get("/todo/:id", handlers.GetTodo)
	api.Post("/todos", handlers.CreateTodo)
	api.Put("/todo/:id", handlers.UpdateTodo)
	api.Delete("todo/:id", handlers.DeleteTodo)
}
