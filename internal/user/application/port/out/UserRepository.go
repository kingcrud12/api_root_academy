package out

import "root_academy_go/internal/user/domain/model"

type IUserRepository interface {
	SaveUser(user *model.User) error
	GetUserByEmail(email string) (*model.User, error)
	GetUserByID(id string) (*model.User, error)
	UpdateUser(id string) (*model.User, error)
	DeleteUser(id string) error
	GetUsers() []model.User
}
