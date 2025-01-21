package UserRepository

import (
	"errors"
	UserDTOS "main/cmd/internals/dtos"
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
func (userRepo *UserRepository) CreateUser(newUser *UserDTOS.CreateUserDTO) (err error) {

	var existingUser UserModel.UserModel
	if err := userRepo.DB.Where("username=?", newUser.Username).First(&existingUser).Error; err == nil {
		return errors.New("username already exists")
	}

	if err := userRepo.DB.Where("email=?", newUser.Email).First(&existingUser).Error; err == nil {
		return errors.New("email already taken")
	}
	userDetails := UserModel.UserDetailModel{
		Address:       newUser.UserDetails.Address,
		ContactNumber: newUser.UserDetails.ContactNumber,
	}
	user := UserModel.UserModel{
		Username:    newUser.Username,
		Password:    newUser.Password,
		Email:       newUser.Email,
		UserDetails: &userDetails,
	}
	userRepo.DB.Create(&user)

	return nil
}
