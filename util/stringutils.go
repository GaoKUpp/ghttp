package util

import "strings"

func IsEmptyString(v string) bool {
	return len(v) == 0
}

func EqualString(s, v string) bool {
	result := false

	if strings.Compare(s, v) == 0 {
		result = true
	}

	return result
}
