package enforcements

import (
	"fmt"
	"strconv"
)


func HandleMaxStr(fieldValue, fieldName, opt string) string {
	opt = ExtractNumber(opt)
	
	minVal, err := strconv.Atoi(opt)
	if err != nil {
		return fmt.Sprintf("Invalid max value for field '%s'", fieldName)
	}

	if len(fieldValue) > minVal {
		return fmt.Sprintf("Field '%s' must be at most %d characters long", fieldName, minVal)
	}
	return ""
}

func HandleMaxInt(fieldValue int64, fieldName, opt string) string {
	opt = ExtractNumber(opt)
	
	minVal, err := strconv.Atoi(opt)
	if err != nil {
		return fmt.Sprintf("Invalid max value for field '%s'", fieldName)
	}

	if int(fieldValue) > minVal {
		return fmt.Sprintf("Field '%s' must be at most %d", fieldName, minVal)
	}

	return ""
}
