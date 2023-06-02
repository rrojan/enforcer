package enforcements

import (
	"fmt"
	"regexp"
	"strings"
)

func matchPattern(pattern, fieldValue, fieldName, customErrorMessage string) string {
	match, err := regexp.MatchString(pattern, fieldValue)
	if err != nil {
		return fmt.Sprintf("Invalid pattern for field '%s'", fieldName)
	}
	if !match {
		if customErrorMessage != "" {
			return customErrorMessage
		}
		return fmt.Sprintf("Field '%s' does not match email pattern", fieldName)
	}
	return ""
}

func HandleMatch(fieldValue, fieldName, opt string) string {
	pattern := ""
	customErrorMessage := ""
	switch opt {
	case "match:email":
		pattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	case "match:phone":
		pattern = `^[0-9\\-]{7,12}$`
	case "match:password":
		// At least one uppercase letter, one lowercase letter,
		// one digit, and one special character
		pattern = `^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*\p{P})(?!.*\s).*$`
		customErrorMessage = "Passwords must contain at least one uppercase letter, one lowercase letter, one digit, and one special character"
	default:
		pattern = strings.TrimPrefix(opt, "match:")
	}

	return matchPattern(pattern, fieldValue, fieldName, customErrorMessage)
}
