package service

import (
	"errors"
	"fmt"

	"root_academy_go/internal/domain/model/user"
	"root_academy_go/internal/domain/repository"
	"root_academy_go/internal/http/dto"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var ErrEmailAlreadyExists = errors.New("email already exists")

func EmailExists(repo repository.UserRepository, email string) (bool, error) {
	u, err := repo.GetUserByEmail(email)
	if err != nil {
		return false, err
	}
	return u != nil, nil
}

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func CreateUser(repo repository.UserRepository, req dto.CreateUserInput) (*dto.UserCreated, error) {
	exists, err := EmailExists(repo, req.Email)
	if err != nil {
		return nil, fmt.Errorf("check email exists: %w", err)
	}
	if exists {
		return nil, ErrEmailAlreadyExists
	}

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	newUser := &user.User{
		ID:        uuid.NewString(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashedPassword,
	}

	if err := repo.CreateUser(newUser); err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return &dto.UserCreated{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
	}, nil
}
