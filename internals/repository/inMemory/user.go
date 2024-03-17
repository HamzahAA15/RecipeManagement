package repository

import (
	"fmt"
	"recipeApps/internals/models"
	"recipeApps/internals/repository"
)

var MapUser = map[string]*models.User{
	"hamzah": {
		ID:             "1",
		Name:           "hamzah",
		Email:          "alfauzi.hamzah@gmail.com",
		HashedPassword: "$2a$12$VchXDt2RSWzgnKKDWXSwzucEGTp6dAcuhShYJdRVHa64csDgVD6Ti",
	},
}

type UserInMemory struct {
}

func NewUserInMemoryRepository() repository.IUserRepository {
	return &UserInMemory{}
}

func (u *UserInMemory) GetUser(name string) (*models.User, error) {
	//as if retrieve from database
	user, ok := MapUser[name]
	if !ok {
		return nil, fmt.Errorf("user %s not found", name)
	}

	return user, nil
}

func (u *UserInMemory) CreateUser(user *models.User) error {
	//as if insert to database
	MapUser[user.Name] = user
	return nil
}
