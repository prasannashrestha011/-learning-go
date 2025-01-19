package userrepository

import (
	"errors"
	"log"
	"main/cmd/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func InitRepo(DB *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: DB,
	}
}
func (userRepository *UserRepository) CreateUser(newUser *models.UserModel) error {
	if newUser == nil {
		return errors.New("insufficent details")
	}
	err := userRepository.DB.Create(newUser).Error
	if err != nil {
		log.Println("Error->", err.Error())
		return err
	}
	return nil
}
