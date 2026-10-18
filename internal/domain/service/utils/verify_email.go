package utils

import (
	"root_academy_go/internal/domain/repository"
)

func EmailAlreadyExists(repo repository.UserRepository, email string) bool {
	return EmailExists(repo, email)
}
