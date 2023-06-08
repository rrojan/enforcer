package enforcements

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func HandleExcludeStr(fieldValue, fieldName, opt string) string {
	excludeValues := strings.Split(strings.TrimPrefix(opt, "exclude:"), ",")
	for _, exclude := range excludeValues {
		if fieldValue == exclude {
			return fmt.Sprintf("Field '%s' contains excluded value: %s", fieldName, exclude)
		}
	}
	return ""
}

func HandleExcludeIntOrFloat(value interface{}, fieldName string, enumOptions string) string {
	excludeValues := strings.Split(strings.TrimPrefix(enumOptions, "exclude:"), ",")
	for _, enumStr := range excludeValues {
		switch reflect.TypeOf(value).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			exclude, err := strconv.ParseInt(enumStr, 10, 64)
			if err != nil {
				return fmt.Sprintf("Invalid exclude value '%s' for field '%s'", enumStr, fieldName)
			}
			if reflect.ValueOf(value).Int() == exclude {
				return fmt.Sprintf("Field '%s' contains excluded value: %d", fieldName, exclude)
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			exclude, err := strconv.ParseUint(enumStr, 10, 64)
			if err != nil {
				return fmt.Sprintf("Invalid exclude value '%s' for field '%s'", enumStr, fieldName)
			}
			if reflect.ValueOf(value).Uint() == exclude {
				return fmt.Sprintf("Field '%s' contains excluded value: %d", fieldName, exclude)
			}
		case reflect.Float32, reflect.Float64:
			exclude, err := strconv.ParseFloat(enumStr, 64)
			if err != nil {
				return fmt.Sprintf("Invalid exclude value '%s' for field '%s'", enumStr, fieldName)
			}
			if reflect.ValueOf(value).Float() == exclude {
				return fmt.Sprintf("Field '%s' contains excluded value: %f", fieldName, exclude)
			}
		default:
			return fmt.Sprintf("Unsupported type for field '%s'", fieldName)
		}
	}

	return ""
}
