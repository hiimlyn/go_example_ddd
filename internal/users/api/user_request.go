package api

// UserCreate is the struct for creating a new user
// validate:
// - username: required, min=3, max=20, unique=users.username (check if username is already taken)
// - password: required, min=8, max=20
// - email: required, email, unique=users.email (check if email is already taken)
type UserCreate struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=8,max=20"`
	Email    string `json:"email" validate:"required,email"`
}
