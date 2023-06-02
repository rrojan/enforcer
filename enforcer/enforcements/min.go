package enforcements

import (
	"fmt"
	"reflect"
	"strconv"
)

func HandleMin(fieldValue, fieldName, opt string) string {
	switch reflect.ValueOf(fieldValue).Kind() {
	case reflect.String:
		minVal, err := strconv.Atoi(opt)
		fmt.Printf("\nString%+v\n", minVal)
		if err != nil {
			return fmt.Sprintf("Invalid minimum value for field '%s'", fieldName)
		}

		if len(fieldValue) < minVal {
			return fmt.Sprintf("Field '%s' must be at least %d characters long", fieldName, minVal)
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fieldInt, err := strconv.Atoi(fieldValue)
		fmt.Printf("\nInt%+v\n", fieldInt)
		if err != nil {
			return fmt.Sprintf("Field '%s' must be a numeric value", fieldName)
		}

		minVal, err := strconv.Atoi(opt)
		if err != nil {
			return fmt.Sprintf("Invalid minimum value for field '%s'", fieldName)
		}

		if fieldInt < minVal {
			return fmt.Sprintf("Field '%s' must be at least %d", fieldName, minVal)
		}

	default:
		return fmt.Sprintf("Unsupported type for field '%s'", fieldName)
	}

	return ""
}
