package utils

import (
	"regexp"
)

var emailFormatRegex = regexp.MustCompile(`^[a-z]+@[a-z]+\.[a-z]{2,}$`)

func ValidEmail(email string) bool {
	return emailFormatRegex.MatchString(email)
}
