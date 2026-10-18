package utils

import (
	"errors"
	"root_academy_go/internal/domain/repository"
)

var ErrEmailAlreadyExists = errors.New("email already exists")

func EmailExists(repo repository.UserRepository, email string) bool {
	_, err := repo.GetUserByEmail(email)
	if err != nil {
		return false
	}
	return true
}
