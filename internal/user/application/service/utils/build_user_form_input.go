package utils

import (
	"root_academy_go/internal/user/adapter/in/web/dto"
	"root_academy_go/internal/user/domain/model"

	"github.com/google/uuid"
)

func BuildUserFromInput(input *dto.Register) *model.User {
	return &model.User{
		ID:        uuid.NewString(),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
	}
}
