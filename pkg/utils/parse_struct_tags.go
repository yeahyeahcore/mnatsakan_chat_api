package utils

import (
	"reflect"
)

// ParseStructTags returns array of struct's tags
func ParseStructTags(object interface{}, tagName string) []string {
	fieldsCount := reflect.TypeOf(object).NumField()
	values := make([]string, fieldsCount)

	for index := 0; index < fieldsCount; index++ {
		field := reflect.TypeOf(object).Field(index)

		values[index] = string(field.Tag.Get(tagName))
	}

	return values
}
