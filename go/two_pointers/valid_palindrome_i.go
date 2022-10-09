package two_pointers

import (
	"strings"
)

func IsValidPalindrome(s string) bool {
	s = strip(s)
	s = strings.ToLower(s)
	if s == "" {
		return true
	}

	arr := strings.Split(s, "")

	last := len(arr) - 1

	for i := 0; i < last; i++ {
		if arr[i] != arr[last] {
			return false
		}

		if i == last {
			break
		}

		last--
	}

	return true
}

func strip(s string) string {
	var result strings.Builder

	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') {
			result.WriteByte(b)
		}
	}
	return result.String()
}
