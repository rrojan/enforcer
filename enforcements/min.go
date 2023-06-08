package enforcements

import (
	"fmt"
	"strconv"
)

func HandleMinStr(fieldValue, fieldName, opt string) string {
	opt = ExtractNumber(opt)

	minVal, err := strconv.Atoi(opt)
	if err != nil {
		return fmt.Sprintf("Invalid minimum value for field '%s'", fieldName)
	}

	if len(fieldValue) < minVal {
		return fmt.Sprintf("Field '%s' must be at least %d characters long", fieldName, minVal)
	}
	return ""
}

func HandleMinInt(fieldValue int64, fieldName, opt string) string {
	opt = ExtractNumber(opt)

	minVal, err := strconv.Atoi(opt)
	if err != nil {
		return fmt.Sprintf("Invalid minimum value for field '%s'", fieldName)
	}

	if int(fieldValue) < minVal {
		return fmt.Sprintf("Field '%s' must be at least %d", fieldName, minVal)
	}

	return ""
}
