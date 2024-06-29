package utils

import "regexp"

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func EmailValidator(email string) bool {

	return emailRegex.MatchString(email)
}