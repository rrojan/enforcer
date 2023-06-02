package enforcements

import (
	"fmt"
	"reflect"
	"strconv"
)

func HandleMax(fieldValue, fieldName, opt string) string {
	switch reflect.ValueOf(fieldValue).Kind() {
	case reflect.String:
		maxVal, err := strconv.Atoi(opt)
		if err != nil {
			return fmt.Sprintf("Invalid maximum value for field '%s'", fieldName)
		}

		if len(fieldValue) > maxVal {
			return fmt.Sprintf("Field '%s' must be at most %d characters long", fieldName, maxVal)
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fieldInt, err := strconv.Atoi(fieldValue)
		if err != nil {
			return fmt.Sprintf("Field '%s' must be a numeric value", fieldName)
		}

		maxVal, err := strconv.Atoi(opt)
		if err != nil {
			return fmt.Sprintf("Invalid maximum value for field '%s'", fieldName)
		}

		if fieldInt > maxVal {
			return fmt.Sprintf("Field '%s' must be at most %d", fieldName, maxVal)
		}

	default:
		return fmt.Sprintf("Unsupported type for field '%s'", fieldName)
	}

	return ""
}
