package utils

import (
	"regexp"
)

// IsNumber - checking whether a string is a number
func IsNumber(data string) bool {
	match, err := regexp.MatchString("([a-z]+)", data)
	if err != nil || match {
		return false
	}

	return true
}
