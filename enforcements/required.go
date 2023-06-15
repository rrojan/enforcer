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
