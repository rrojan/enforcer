package enforcements

import (
	"fmt"
	"strconv"
	"strings"
)

func HandleBetween(fieldValue, fieldName, opt string) string {
	rangeVals := strings.Split(strings.TrimPrefix(opt, "between:"), ",")
	if len(rangeVals) != 2 {
		return fmt.Sprintf("Invalid range values for field '%s'", fieldName)
	}

	min, err := strconv.Atoi(rangeVals[0])
	if err != nil {
		return fmt.Sprintf("Invalid range values for field '%s'", fieldName)
	}

	max, err := strconv.Atoi(rangeVals[1])
	if err != nil {
		return fmt.Sprintf("Invalid range values for field '%s'", fieldName)
	}

	if len(fieldValue) < min || len(fieldValue) > max {
		return fmt.Sprintf("Field '%s' must be between %s and %s characters", fieldName, rangeVals[0], rangeVals[1])
	}

	return ""
}
