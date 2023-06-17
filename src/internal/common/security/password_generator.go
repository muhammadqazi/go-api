package security

import (
	"crypto/rand"
	"math/big"
)

func getRandomChar(chars string) string {
	index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
	return string(chars[index.Int64()])
}

func GeneratePassword(length, retries int) string {
	lowerChars := "abcdefghijklmnopqrstuvwxyz"
	upperChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars := "0123456789"
	specialChars := "!@#$%^&*"

	password := ""
	types := 0

	// Generate at least one character of each type
	password += getRandomChar(lowerChars)
	password += getRandomChar(upperChars)
	password += getRandomChar(numberChars)
	password += getRandomChar(specialChars)
	types = 15

	// Generate remaining characters randomly
	for i := 4; i < length; i++ {
		randomType, _ := rand.Int(rand.Reader, big.NewInt(4))
		switch randomType.Int64() + 1 {
		case 1:
			password += getRandomChar(lowerChars)
			types |= 1
		case 2:
			password += getRandomChar(upperChars)
			types |= 2
		case 3:
			password += getRandomChar(numberChars)
			types |= 4
		case 4:
			password += getRandomChar(specialChars)
			types |= 8
		}
	}

	if types != 15 && retries < 20 {
		return GeneratePassword(length, retries+1)
	}

	return password
}
