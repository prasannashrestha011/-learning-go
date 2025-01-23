package UserService

//package
import (
	UserDTOS "main/cmd/internals/dtos"
	UserRepository "main/cmd/internals/repositories"
	"main/cmd/internals/services/security"
	"main/cmd/pkgs/dtos"
	schemavalidators "main/cmd/pkgs/schema_validators"
)

type UserService struct {
	UserRepo *UserRepository.UserRepository
}

func InitUserService(userRepo *UserRepository.UserRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (s *UserService) RegisterUser(newUser *UserDTOS.CreateUserDTO) (response dtos.ResponseDto) {
	isEmailValid, err := schemavalidators.ValidateEmail(newUser.Email)
	if !isEmailValid {

		return dtos.ResponseDto{
			Success: false,
			Message: err.Error(),
		}
	}
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
func (s *UserService) AuthenticatedUser(user UserDTOS.AuthUserDTO) *dtos.DataDto {
	authUser, err := s.UserRepo.FetchUserByUsername(user.Username)
	if err != nil {
		return &dtos.DataDto{
			Data:       nil,
			StatusCode: 403,
			Success:    false,
		}
	}
	isAuthenticated := security.ComparePassword(authUser.Password, user.Password)
	if !isAuthenticated {
		errMessage := "Invalid password, please try again"
		return &dtos.DataDto{
			Data:       nil,
			StatusCode: 403,
			Success:    false,
			Message:    &errMessage,
		}
	}
	return &dtos.DataDto{
		Data:       authUser,
		StatusCode: 200,
		Success:    true,
	}
}
func (s *UserService) GetUserByID(userId string) dtos.ResponseDto {
	if userId == "" {
		return dtos.ResponseDto{
			Message: "Please provided valid user id",
			Success: false,
		}
	}
	foundUser, err := s.UserRepo.FetchUserByID(userId)
	if err != nil {
		return dtos.ResponseDto{
			Message: err.Error(),
			Success: false,
		}
	}
	return dtos.ResponseDto{
		Message:    "Found User",
		StatusCode: 200,
		Success:    true,
		Data:       foundUser,
	}
}
func (s *UserService) GetAllUsers() dtos.ResponseDto {
	userList, err := s.UserRepo.FetchAllUsers()
	if err != nil {
		return dtos.ResponseDto{
			Message:    err.Error(),
			Success:    false,
			StatusCode: 404,
		}
	}
	return dtos.ResponseDto{
		Message:    "User list",
		Success:    true,
		Data:       userList,
		StatusCode: 200,
	}
}
