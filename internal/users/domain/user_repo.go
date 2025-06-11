package domain

type UserRepo interface {
	GetUser(id string) (UserEntity, error)
	GetAllUsers() ([]UserEntity, error)
	CreateUser(user UserEntity) error
}
