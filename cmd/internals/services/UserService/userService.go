package services

import (
	userrepository "main/cmd/internals/repositories/UserRepository"
	"main/cmd/models"
)

type UserService struct {
	userRepository *userrepository.UserRepository
}

func InitService(userRepo *userrepository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepo,
	}
}
func (userService *UserService) RegisterNewUser(newUser *models.UserModel) error {
	err := userService.userRepository.CreateUser(newUser)
	if err != nil {
		return err
	}
	return nil
}
