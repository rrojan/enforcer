package enforcements

import (
	"fmt"
	"regexp"
	"strings"
)

func HandleMatch(fieldValue, fieldName, opt string) string {
	if opt == "match:email" {
		emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		match, err := regexp.MatchString(emailPattern, fieldValue)
		if err != nil {
			return fmt.Sprintf("Invalid pattern for field '%s'", fieldName)
		}

		if !match {
			return fmt.Sprintf("Field '%s' does not match email pattern", fieldName)
		}

		return ""
	}

	pattern := strings.TrimPrefix(opt, "match:")
	match, err := regexp.MatchString(pattern, fieldValue)
	if err != nil {
		return fmt.Sprintf("Invalid pattern for field '%s'", fieldName)
	}

	if !match {
		return fmt.Sprintf("Field '%s' does not match pattern", fieldName)
	}

	return ""
}
