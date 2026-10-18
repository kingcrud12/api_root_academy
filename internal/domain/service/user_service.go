package service

import (
	"fmt"

	"root_academy_go/internal/domain/repository"
	"root_academy_go/internal/domain/service/utils"
	"root_academy_go/internal/http/dto"
)

func CreateUser(repo repository.UserRepository, req dto.CreateUserInput) (*dto.UserCreated, error) {
	if utils.EmailAlreadyExists(repo, req.Email) {
		return nil, utils.ErrEmailAlreadyExists
	}

	if !utils.ValidEmail(req.Email) {
		return nil, fmt.Errorf("Email au format invalide")
	}

	newUser, err := utils.BuildUserFromInput(req)
	if err != nil {
		return nil, err
	}

	if err := repo.SaveUser(newUser); err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return utils.FormatUser(newUser), nil
}
