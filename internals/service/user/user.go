package user

import (
	"recipeApps/internals/models"
	"recipeApps/internals/repository"
	"recipeApps/utils/middleware/authentication"
	"recipeApps/utils/request"
	"recipeApps/utils/response"

	"github.com/google/uuid"
)

type userService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(userReq request.CreateUser) error {
	hashedPass, err := authentication.GenerateHashedPassword(userReq.Password)
	if err != nil {
		return err
	}

	user := &models.User{
		ID:             uuid.NewString(),
		Name:           userReq.Name,
		Email:          userReq.Email,
		HashedPassword: hashedPass,
	}

	err = s.repo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) GetUserByName(name string) (response.UserResponseData, error) {
	user, err := s.repo.GetUser(name)
	if err != nil {
		return response.UserResponseData{}, err
	}

	return response.UserResponseData{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
