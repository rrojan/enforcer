package enforcements

import (
	"fmt"
	"strconv"
	"strings"
)

func HandleWordCount(fieldValue, fieldName, opt string) string {
	rangeVals := strings.Split(strings.TrimPrefix(opt, "wordCount:"), ",")
	fmt.Printf("\n%#v\n", rangeVals)

	if len(rangeVals) != 2 {
		return fmt.Sprintf("Invalid word count range for field '%s'", fieldName)
	}

	min, err := strconv.Atoi(rangeVals[0])
	if err != nil {
		return fmt.Sprintf("Invalid word count range for field '%s'", fieldName)
	}

	max, err := strconv.Atoi(rangeVals[1])
	if err != nil {
		return fmt.Sprintf("Invalid word count range for field '%s'", fieldName)
	}

	words := countWords(fieldValue)
	if words < min || words > max {
		return fmt.Sprintf("Field '%s' must be between %s and %s words", fieldName, rangeVals[0], rangeVals[1])
	}

	return ""
}
