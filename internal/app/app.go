package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hiimlyn/go-api/internal/commons"
	"github.com/hiimlyn/go-api/internal/config"
	"github.com/hiimlyn/go-api/internal/users"
	"github.com/hiimlyn/go-api/internal/users/domain"
)

func New(cfg *config.Config) *fiber.App {
	app := fiber.New()

	//load middleware
	app.Use(logger.New())

	//load database
	db := commons.ConnectDB(cfg)
	db.AutoMigrate(&domain.UserEntity{})

	// Dependency Injection
	users.Inject(app, db)

	return app
}
