package security

import (
	"root_academy_go/internal/user/application/port/out"

	"golang.org/x/crypto/bcrypt"
)

type BcryptPasswordHasher struct{}

var _ out.IPasswordHasher = (*BcryptPasswordHasher)(nil)

func NewBcryptPasswordHasher() *BcryptPasswordHasher {
	return &BcryptPasswordHasher{}
}

func (h *BcryptPasswordHasher) HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}
