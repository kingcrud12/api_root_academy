package utils

import "root_academy_go/internal/user/domain/rules"

func EmailValid(email string) bool {
	IsEmailValid := rules.ValidEmail(email)
	if !IsEmailValid {
		return false
	}
	return true
}
