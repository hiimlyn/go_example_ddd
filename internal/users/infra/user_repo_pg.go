package infra

import (
	"log"

	"github.com/hiimlyn/go-api/internal/users/domain"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) domain.UserRepo {
	return &userRepo{db: db}
}

func (repo *userRepo) CreateUser(user domain.UserEntity) error {
	if err := repo.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *userRepo) GetUser(id string) (domain.UserEntity, error) {
	var user domain.UserEntity
	if err := repo.db.Where("id = ?", id).First(&user).Error; err != nil {
		return domain.UserEntity{}, err
	}
	return user, nil
}

func (repo *userRepo) GetAllUsers() ([]domain.UserEntity, error) {
	var users []domain.UserEntity
	log.Println("GetAllUsers")
	if err := repo.db.Find(&users).Error; err != nil {
		return []domain.UserEntity{}, err
	}
	return users, nil
}
