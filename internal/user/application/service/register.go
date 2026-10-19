package service

import (
	"errors"
	"fmt"
	"root_academy_go/internal/user/adapter/in/web/dto"
	"root_academy_go/internal/user/application/port/out"
	"root_academy_go/internal/user/application/service/utils"
)

var email_invalid_error = errors.New("Invalide email format")

var email_already_exists_error = errors.New("email already exists")

func Register(input *dto.Register, repo out.IUserRepository) error {
	if !utils.EmailValid(input.Email) {
		return email_invalid_error
	}
	if utils.EmailExists(input.Email, repo) {
		return email_already_exists_error
	}
	user := utils.BuildUserFromInput(input)

	userSaved := repo.SaveUser(user)

	if userSaved != nil {
		return fmt.Errorf("erreur lors de l'inscription en base")
	}

	return nil

}
