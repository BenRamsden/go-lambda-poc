package model

import (
	"context"
	"errors"
)

type userService struct {
	repo UserRepository
}

// CreateUser implements UserService.
func (svc *userService) CreateUser(context context.Context, newUser User) (User, error) {
	if len(newUser.ID) == 0 {
		return User{}, NewInvalidInputError(errors.New("id is required"))
	}

	if len(newUser.Name) == 0 {
		return User{}, NewInvalidInputError(errors.New("name is required"))
	}

	return svc.repo.CreateUser(context, newUser)
}

// GetUser implements UserService.
func (svc *userService) GetUser(context context.Context, id string) (User, error) {
	if len(id) == 0 {
		return User{}, NewInvalidInputError(errors.New("id is required"))
	}

	return svc.repo.GetUser(context, id)
}

// GetUsers implements UserService.
func (svc *userService) GetUsers(context context.Context, pagination Pagination) ([]User, error) {
	return svc.repo.GetUsers(context, pagination)
}

func NewUserService(repo UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}
