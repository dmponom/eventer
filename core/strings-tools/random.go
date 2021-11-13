package stringtools

import (
	numberTools "eventer/core/numbers-tools"
	"fmt"
)

func GetRandom(len int) string {
	bytes := make([]byte, len)

	for i := 0; i < len; i++ {
		bytes[i] = byte(numberTools.GetRandomIntInRange(asciiFirstLowerCaseLetterCode, asciiLastLowerCaseLetterCode))
	}

	return string(bytes)
}

func GetRandomEmail() string {
	return fmt.Sprintf("%s@%s.%s", GetRandom(10), GetRandom(5), GetRandom(3))
}

func GetRandomEmails(len int) []string {
	emails := make([]string, len)

	for i := 0; i < len; i++ {
		emails[i] = GetRandomEmail()
	}

	return emails
}
