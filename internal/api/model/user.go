package model

import "context"

type User struct {
	ID    string
	Name  string
	Email string
}

type UserRepository interface {
	CreateUser(context context.Context, newUser User) (User, error)
	GetUser(context context.Context, id string) (User, error)
	GetUsers(context context.Context, pagination Pagination) ([]User, error)
}

type UserService interface {
	CreateUser(context context.Context, newUser User) (User, error)
	GetUser(context context.Context, id string) (User, error)
	GetUsers(context context.Context, pagination Pagination) ([]User, error)
}
