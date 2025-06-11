package app

import (
	"github.com/hiimlyn/go-api/internal/users/domain"
	"github.com/hiimlyn/go-api/internal/users/infra"
	"gorm.io/gorm"
)

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

type UserService struct {
	db *gorm.DB
}

func (s *UserService) GetUser(id string) (domain.UserEntity, error) {
	return infra.NewUserRepo(s.db).GetUser(id)
}

func (s *UserService) CreateUser(user domain.UserEntity) error {
	return infra.NewUserRepo(s.db).CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]domain.UserEntity, error) {
	return infra.NewUserRepo(s.db).GetAllUsers()
}
