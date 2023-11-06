package random

import (
	"math/rand"
	"time"
)

const (
	lowerLetters    = "abcdefghijklmnopqrstuvwxyz"
	lowerLen        = len(lowerLetters)
	upperLetters    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	upperLen        = len(upperLetters)
	numbers         = "0123456789"
	numbersLen      = len(numbers)
	specialChars    = "!@#$%^&*()-_+="
	specialCharsLen = len(specialChars)
	minLength       = 8
)

// generateString generates a random string of the given length.
func generateString(length int) string {
	var s []byte
	if length < minLength {
		length = minLength
	}
	for i := 0; i < length; i++ {
		switch i % 4 {
		case 0:
			s = append(s, lowerLetters[rand.Intn(lowerLen)])
		case 1:
			s = append(s, upperLetters[rand.Intn(upperLen)])
		case 2:
			s = append(s, numbers[rand.Intn(numbersLen)])
		default:
			s = append(s, specialChars[rand.Intn(specialCharsLen)])
		}
	}
	return string(s)
}

// GeneratePassword generates a random password of the given length.
func GeneratePassword(length int) string {
	rand.Seed(time.Now().UnixNano())
	password := generateString(length)
	// Generated password
	return password
}

// GenRandString generates a random string of the given length using lowercase letters and numbers.
func GenRandString(length int) string {
	rand.Seed(time.Now().UnixNano())

	charSet := lowerLetters + numbers
	str := make([]byte, length)
	for i := 0; i < length; i++ {
		str[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(str)
}
