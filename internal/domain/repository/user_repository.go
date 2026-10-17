package repository

import "root_academy_go/internal/domain/model/user"

type UserRepository interface {
	CreateUser(user *user.User) error
	GetUserByEmail(email string) (*user.User, error)
	GetUserByID(id string) (*user.User, error)
	UpdateUser(user *user.User) error
	DeleteUser(id string) error
	GetUsers() []user.User
}
