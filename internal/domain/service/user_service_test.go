package service

import (
	"testing"

	"regexp"

	"root_academy_go/internal/domain/model/user"

	"github.com/stretchr/testify/assert"
)

type FakeUserRepository interface {
	Save(u *user.User) error
}

type userRepository struct {
	userCreated bool
	savedUsers  []*user.User
}

func (f *userRepository) Save(u *user.User) error {
	f.userCreated = true
	f.savedUsers = append(f.savedUsers, u)
	return nil
}

var _ FakeUserRepository = (*userRepository)(nil)

var usersTable = []user.User{
	{
		FirstName: "Ada",
		LastName:  "Lovelace",
		Email:     "ada@example.com",
		Password:  "s3cr3t",
	},
	{
		FirstName: "Grace",
		LastName:  "Hopper",
		Email:     "grace@example.com",
		Password:  "an0ther",
	},
}

func TestUserNotExist(t *testing.T) {
	userExists := user.User{
		FirstName: "Ada",
		LastName:  "Lovelace",
		Email:     "ada@example.com",
		Password:  "s3cr3t",
	}

	for _, u := range usersTable {
		assert.NotEqual(t, userExists.Email, u.Email)
	}
}

func TestEmailFormat(t *testing.T) {
	var emailRegex = regexp.MustCompile((`^[a-z]+@[a-z]+\.[a-z]{2,}$`))

	userEmailWrongFormat := "dodo.gmail.com"

	assert.True(t, emailRegex.MatchString(userEmailWrongFormat))

}
