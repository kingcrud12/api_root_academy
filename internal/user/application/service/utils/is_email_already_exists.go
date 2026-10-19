package utils

import "root_academy_go/internal/user/application/port/out"

func EmailExists(email string, repo out.IUserRepository) bool {
	user, err := repo.GetUserByEmail(email)
	if err != nil {
		return false
	}

	return user != nil

}
