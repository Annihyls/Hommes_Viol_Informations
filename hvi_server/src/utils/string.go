/*
Package utils implement utilities like string functions
*/
package utils

import "strings"

func IsStringEmpty(str string) bool {
	return len(strings.TrimSpace(str)) <= 0
}

func StringCut(str string, token string) (string, string, bool) {
	return strings.Cut(str, token)
}

func StringTrimSpace(str string) string {
	return strings.TrimSpace(str)
}

func StringHasPrefix(str string, prefix string) bool {
	return strings.HasPrefix(str, prefix)
}
