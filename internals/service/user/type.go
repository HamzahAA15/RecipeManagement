package user

import (
	"recipeApps/utils/request"
	"recipeApps/utils/response"
)

type IUserService interface {
	GetUserByName(name string) (response.UserResponseData, error)
	CreateUser(userReq request.CreateUser) error
}
