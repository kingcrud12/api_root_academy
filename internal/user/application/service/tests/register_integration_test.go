package tests

import (
	"errors"
	"testing"

	"root_academy_go/internal/user/adapter/in/web/dto"
	"root_academy_go/internal/user/application/service"
	"root_academy_go/internal/user/domain/model"

	"github.com/stretchr/testify/assert"
)

type FakeUserRepository struct {
	usersByEmail map[string]*model.User
}

func NewFakeUserRepository() *FakeUserRepository {
	return &FakeUserRepository{usersByEmail: make(map[string]*model.User)}
}

func (r *FakeUserRepository) SaveUser(user *model.User) error {
	r.usersByEmail[user.Email] = user
	return nil
}

func (r *FakeUserRepository) GetUserByEmail(email string) (*model.User, error) {
	user, exists := r.usersByEmail[email]
	if !exists {
		return nil, nil
	}
	return user, nil
}

func (r *FakeUserRepository) GetUserByID(id string) (*model.User, error) {
	for _, user := range r.usersByEmail {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, nil
}

func (r *FakeUserRepository) UpdateUser(id string) (*model.User, error) {
	return nil, errors.New("not implemented")
}

func (r *FakeUserRepository) DeleteUser(id string) error {
	for email, user := range r.usersByEmail {
		if user.ID == id {
			delete(r.usersByEmail, email)
			return nil
		}
	}
	return nil
}

func (r *FakeUserRepository) GetUsers() []model.User {
	users := make([]model.User, 0, len(r.usersByEmail))
	for _, user := range r.usersByEmail {
		users = append(users, *user)
	}
	return users
}

type FakePasswordHasher struct{}

func (h *FakePasswordHasher) HashPassword(password string) (string, error) {
	return "hashed-" + password, nil
}

func Test_register_creates_user_with_hashed_password(t *testing.T) {
	repo := NewFakeUserRepository()
	userService := service.NewUserService(repo, &FakePasswordHasher{})

	input := &dto.Register{
		FirstName: "Ada",
		LastName:  "Lovelace",
		Email:     "ada@example.com",
		Password:  "secret",
	}

	err := userService.Register(input)

	assert.NoError(t, err)

	saved, _ := repo.GetUserByEmail("ada@example.com")
	assert.NotNil(t, saved)
	assert.Equal(t, "hashed-secret", saved.Password)
}

func Test_register_rejects_invalid_email(t *testing.T) {
	repo := NewFakeUserRepository()
	userService := service.NewUserService(repo, &FakePasswordHasher{})

	input := &dto.Register{Email: "not-an-email", Password: "secret"}

	err := userService.Register(input)

	assert.Error(t, err)
	assert.Empty(t, repo.GetUsers())
}

func Test_register_rejects_duplicate_email(t *testing.T) {
	repo := NewFakeUserRepository()
	userService := service.NewUserService(repo, &FakePasswordHasher{})

	input := &dto.Register{Email: "ada@example.com", Password: "secret"}

	err := userService.Register(input)
	assert.NoError(t, err)

	err = userService.Register(input)
	assert.Error(t, err)
}
