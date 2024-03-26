package password

import (
	"math/rand"
)

const (
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits           = "0123456789"
	specialChars     = "!@#$%^&*()-_=+[]{}|;:'\"<>,.?/~"
)

func GeneratePassword(length int, useUppercase, useDigits, useSpecialChars bool) string {
	var validChars string
	validChars += lowercaseLetters

	if useUppercase {
		validChars += uppercaseLetters
	}
	if useDigits {
		validChars += digits
	}
	if useSpecialChars {
		validChars += specialChars
	}
	if validChars == "" {
		validChars = lowercaseLetters + uppercaseLetters + digits + specialChars
	}

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		password[i] = validChars[rand.Intn(len(validChars))]
	}
	return string(password)
}
