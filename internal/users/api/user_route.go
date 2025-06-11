package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitRoutes(app *fiber.App, db *gorm.DB) {
	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// User group
	userHandler := NewUserHandler(db)
	user := app.Group("/user")
	user.Get("/users", userHandler.GetAllUsers)
	user.Post("/", userHandler.CreateUser)
	user.Get("/:id", userHandler.GetUser)

	// Add more groups/routes here
}
