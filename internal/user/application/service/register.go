package service

import (
	"errors"
	"fmt"
	"root_academy_go/internal/user/adapter/in/web/dto"
	"root_academy_go/internal/user/application/port/in"
	"root_academy_go/internal/user/application/port/out"
	"root_academy_go/internal/user/application/service/utils"
)

var email_invalid_error = errors.New("Invalide email format")

var email_already_exists_error = errors.New("email already exists")

type UserService struct {
	repo   out.IUserRepository
	hasher out.IPasswordHasher
}

var _ in.RegisterUser = (*UserService)(nil)

func NewUserService(repo out.IUserRepository, hasher out.IPasswordHasher) *UserService {
	return &UserService{repo: repo, hasher: hasher}
}

func (s *UserService) Register(input *dto.Register) error {
	if !utils.EmailValid(input.Email) {
		return email_invalid_error
	}
	if utils.EmailExists(input.Email, s.repo) {
		return email_already_exists_error
	}
	user := utils.BuildUserFromInput(input)

	hashedPassword, err := s.hasher.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("erreur lors du hachage du mot de passe")
	}
	user.Password = hashedPassword

	if err := s.repo.SaveUser(user); err != nil {
		return fmt.Errorf("erreur lors de l'inscription en base")
	}

	return nil
}
