package enforcements

import (
	"fmt"
)

func HandleRequired(fieldValue, fieldName string) string {
	if fieldValue == "" {
		return fmt.Sprintf("Required field '%s' is not provided", fieldName)
	}

	return ""
}
