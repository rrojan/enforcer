package enforcements

import (
	"fmt"
	"reflect"
)

func HandleRequired(fieldValue reflect.Value, fieldName string) string {
	if IsEmpty(fieldValue)  {
		return fmt.Sprintf("Required field '%s' is not provided", fieldName)
	}

	return ""
}

func HandleRequiredAbsent(fieldValue reflect.Value, fieldName string) string {
	if !IsEmpty(fieldValue) {
		return fmt.Sprintf("Field %s cannot be set to a value", fieldName)
	}
	return ""
}
