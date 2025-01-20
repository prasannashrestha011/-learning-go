package UserRepository

import (
	"errors"
	UserModel "main/cmd/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func InitUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}
func (userRepo *UserRepository) CreateUser(newUser *UserModel.UserModel) (err error) {

	if newUser == nil {
		return errors.New("insufficent details for registration")
	}
	var existingUser UserModel.UserModel
	if err := userRepo.DB.Where("username=?", newUser.Username).First(&existingUser).Error; err == nil {
		return errors.New("username already exists")
	}

	if err := userRepo.DB.Where("email=?", newUser.Email).First(&existingUser).Error; err == nil {
		return errors.New("email already taken")
	}
	userRepo.DB.Create(&newUser)
	return nil
}
