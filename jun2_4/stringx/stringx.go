package stringx

import (
	"slices"
)

func Reverse(s string) string {
	r := []rune(s)
	slices.Reverse(r)

	return string(r)
}

func IsEmpty(s string) bool {
	return s == ""
}
