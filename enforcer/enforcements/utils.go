package enforcements

import (
	"regexp"
	"strings"
)


func ExtractNumber(str string) string {
	re := regexp.MustCompile(`\d+`)
	match := re.FindString(str)
	return match
}

func containsUppercase(s string) bool {
	for _, c := range s {
		if 'A' <= c && c <= 'Z' {
			return true
		}
	}
	return false
}

func containsLowercase(s string) bool {
	for _, c := range s {
		if 'a' <= c && c <= 'z' {
			return true
		}
	}
	return false
}

func containsDigit(s string) bool {
	for _, c := range s {
		if '0' <= c && c <= '9' {
			return true
		}
	}
	return false
}

func containsSpecialCharacter(s string) bool {
	specialChars := `!@#$%^&*()_+-=[]{}|;:'",.<>/?`
	for _, c := range s {
		if strings.ContainsRune(specialChars, c) {
			return true
		}
	}
	return false
}
