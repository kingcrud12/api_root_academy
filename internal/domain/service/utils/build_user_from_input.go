package utils

import (
	"fmt"
	"root_academy_go/internal/domain/model/user"
	"root_academy_go/internal/http/dto"

	"github.com/google/uuid"
)

func BuildUserFromInput(req dto.CreateUserInput) (*user.User, error) {
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	return &user.User{
		ID:        uuid.NewString(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashedPassword,
	}, nil
}
