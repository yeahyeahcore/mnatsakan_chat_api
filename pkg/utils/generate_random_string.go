package utils

import (
	"math/rand"
	"time"
)

// GenerateRandomString - generates random string by length
func GenerateRandomString(length int) string {
	code := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/~!@#$%^&*()_="
	data := make([]byte, length)

	rand.Seed(time.Now().UnixNano())

	for index := 0; index < length; index++ {
		randomСharIndex := rand.Intn(len(code))
		data[index] = code[randomСharIndex]
	}

	return string(data)
}
