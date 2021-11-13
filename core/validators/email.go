package validators

import (
	"net/mail"
	"regexp"
)

func IsEmailAddress(email string) bool {
	if _, err := mail.ParseAddress(email); err != nil {
		return false
	}

	match, err := regexp.MatchString(emailValidateRegex, email)
	if err != nil {
		return false
	}

	return match
}
