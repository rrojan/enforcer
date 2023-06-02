package enforcements

import (
	"fmt"
	"regexp"
	"strings"
)

func matchPattern(pattern, fieldValue, fieldName string) string {
	match, err := regexp.MatchString(pattern, fieldValue)
	if err != nil {
		return fmt.Sprintf("Invalid pattern for field '%s'", fieldName)
	}
	if !match {
		return fmt.Sprintf("Field '%s' does not match email pattern", fieldName)
	}
	return ""
}

func HandleMatch(fieldValue, fieldName, opt string) string {
	var pattern string
	switch opt {
	case "match:email":
		pattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	case "match:phone":
		pattern = `^[0-9\\-]{7,12}$`
	default:
		pattern = strings.TrimPrefix(opt, "match:")
	}

	return matchPattern(pattern, fieldValue, fieldName)
}
