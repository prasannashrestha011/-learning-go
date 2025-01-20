package UserService

import (
	UserRepository "main/cmd/internals/repositories"
	"main/cmd/internals/services/security"
	UserModel "main/cmd/models"
	"main/cmd/pkgs/dtos"
)

type UserService struct {
	UserRepo *UserRepository.UserRepository
}

func InitUserService(userRepo *UserRepository.UserRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (s *UserService) RegisterUser(newUser *UserModel.UserModel) (response dtos.ResponseDto) {
	hashedPassword, err := security.HashPassword(newUser.Password)
	if err != nil {
		return dtos.ResponseDto{
			Message:    err.Error(),
			Success:    false,
			StatusCode: 500,
		}
	}
	newUser.Password = hashedPassword
	err = s.UserRepo.CreateUser(newUser)
	if err != nil {
		return dtos.ResponseDto{
			Message:    err.Error(),
			Success:    false,
			StatusCode: 400,
		}
	}
	return dtos.ResponseDto{
		Message:    "New User created",
		Success:    true,
		StatusCode: 200,
	}
}
