package utils

import (
	"root_academy_go/internal/domain/model/user"
	"root_academy_go/internal/http/dto"
)

func FormatUser(user *user.User) *dto.UserCreated {
	return &dto.UserCreated{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
