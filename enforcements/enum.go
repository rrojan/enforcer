package enforcements

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func HandleEnumStr(fieldValue, fieldName, opt string) string {
	enumValues := strings.Split(strings.TrimPrefix(opt, "enum:"), ",")
	for _, enum := range enumValues {
		if fieldValue == enum {
			return "" // Value is in the enum, no error
		}
	}
	return fmt.Sprintf("Field '%s' does not match any valid enum value", fieldName)
}

func HandleEnumIntOrFloat(value interface{}, fieldName string, enumOptions string) string {
	enumValues := strings.Split(strings.TrimPrefix(enumOptions, "enum:"), ",")
	for _, enumStr := range enumValues {
		switch reflect.TypeOf(value).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			enum, err := strconv.ParseInt(enumStr, 10, 64)
			if err != nil {
				return fmt.Sprintf("Invalid enum value '%s' for field '%s'", enumStr, fieldName)
			}
			if reflect.ValueOf(value).Int() == enum {
				return "" // Value is in the enum, no error
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			enum, err := strconv.ParseUint(enumStr, 10, 64)
			if err != nil {
				return fmt.Sprintf("Invalid enum value '%s' for field '%s'", enumStr, fieldName)
			}
			if reflect.ValueOf(value).Uint() == enum {
				return "" // Value is in the enum, no error
			}
		case reflect.Float32, reflect.Float64:
			enum, err := strconv.ParseFloat(enumStr, 64)
			if err != nil {
				return fmt.Sprintf("Invalid enum value '%s' for field '%s'", enumStr, fieldName)
			}
			if reflect.ValueOf(value).Float() == enum {
				return "" // Value is in the enum, no error
			}
		default:
			return fmt.Sprintf("Unsupported type for field '%s'", fieldName)
		}
	}

	return fmt.Sprintf(
		"Field '%s' does not match any enum values: %s",
		fieldName, strings.Join(enumValues, ", "),
	)
}
