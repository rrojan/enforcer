package enforcements

import (
	"reflect"
	"regexp"
	"strings"
)

func ExtractNumber(str string) string {
	re := regexp.MustCompile(`\d+`)
	match := re.FindString(str)
	return match
}

func IsString(k reflect.Kind) bool {
	return k == reflect.String
}

func IsIntType(k reflect.Kind) bool {
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	default:
		return false
	}
}

func IsFloatType(k reflect.Kind) bool {
	switch k {
	case reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func IsEmpty(v reflect.Value) bool{
	return (IsString(v.Kind()) && v.String() == "") || (IsIntType(v.Kind()) && v.Int() == 0) || (IsFloatType(v.Kind()) && v.Float() == 0.0)
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

func countWords(s string) int {
	spaceCount := 0
	for _, c := range s {
		if c == ' ' {
			spaceCount = spaceCount + 1
		}
	}
	return spaceCount + 1
}

func ArrayContainsSubstr(a []string, s string) bool {
	for _, elem := range a {
		if strings.Contains(elem, s) {
			return true
		}
	}
	return false
}
