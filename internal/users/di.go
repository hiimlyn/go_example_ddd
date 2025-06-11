package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hiimlyn/go-api/internal/users/api"
	"gorm.io/gorm"
)

func Inject(app *fiber.App, db *gorm.DB) {
	api.InitRoutes(app, db)
}
