package repository

import "root_academy_go/internal/domain/model/user"

type UserRepository interface {
	SaveUser(user *user.User) error
	GetUserByEmail(email string) (*user.User, error)
	GetUserByID(id string) (*user.User, error)
	UpdateUser(id string) (*user.User, error)
	DeleteUser(id string) error
	GetUsers() []user.User
}
