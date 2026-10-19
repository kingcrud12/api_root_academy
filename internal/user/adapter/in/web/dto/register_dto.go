package dto

import "root_academy_go/internal/user/domain/model"

type Register struct {
	FirstName string
	LastName  string
	Email     string
	Roles     model.Role
	Password  string
}
