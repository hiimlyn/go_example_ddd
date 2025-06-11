package api

import (
	"encoding/json"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hiimlyn/go-api/internal/commons/model"
	"github.com/hiimlyn/go-api/internal/users/app"
	"github.com/hiimlyn/go-api/internal/users/domain"
	"gorm.io/gorm"
)

var validate = validator.New()

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := app.NewUserService(h.db).GetUser(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}
	// convert user entity to user struct
	userDto := UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return c.JSON(userDto)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {

	var userRequest UserCreate

	if err := c.BodyParser(&userRequest); err != nil {
		return c.Status(400).JSON(model.Response{
			Status:  "error",
			Message: "invalid request",
			Error:   err.Error(),
		})
	}

	if err := validate.Struct(&userRequest); err != nil {
		errs := err.(validator.ValidationErrors)
		messages := make(map[string]string)
		for _, e := range errs {
			messages[e.Field()] = e.Tag()
		}
		//map to string
		json, err := json.Marshal(messages)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to marshal messages"})
		}
		return c.Status(422).JSON(model.Response{
			Status:  "error",
			Message: "validation error",
			Error:   string(json),
		})
	}

	// Validate struct

	user := domain.UserEntity{
		Username:  userRequest.Username,
		Password:  userRequest.Password,
		Email:     userRequest.Email,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	err := app.NewUserService(h.db).CreateUser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to create user"})
	}
	return c.JSON(model.Response{
		Status:  "success",
		Message: "user created successfully",
		Data:    user.ID,
	})
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := app.NewUserService(h.db).GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(model.Response{
			Status:  "error",
			Message: "failed to get all users",
			Error:   err.Error(),
		})
	}
	// convert user entity to user struct
	usersDto := make([]UserResponse, len(users))
	for i, user := range users {
		usersDto[i] = UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}

	return c.JSON(model.Response{
		Status:  "success",
		Message: "users fetched successfully",
		Data:    users,
	})
}
