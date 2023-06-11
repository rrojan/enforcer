package enforcements

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ApplyDefaults(v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("pointer to struct expected, got %T", v)
	}

	// Dereference the pointer to get the struct value
	rv = rv.Elem()

	if rv.Kind() != reflect.Struct {
		return fmt.Errorf("pointer to struct expected, got pointer to %T", v)
	}

	// Iterate over the fields of the struct
	for i := 0; i < rv.NumField(); i++ {
		fieldValue := rv.Field(i)
		fieldType := rv.Type().Field(i)

		// Check if the field has the enforce tag
		tagValue := fieldType.Tag.Get("enforce")
		if tagValue == "" || !strings.Contains(tagValue, "default:") {
			continue
		}

		// Check if the field is empty (zero value)
		if fieldValue.Kind() == reflect.String && fieldValue.String() == "" {
			defaultValue := getDefaultValue(tagValue)
			// Set the default value for the field
			fieldValue.SetString(defaultValue)
		} else if fieldValue.Kind() == reflect.Int || fieldValue.Kind() == reflect.Int8 ||
			fieldValue.Kind() == reflect.Int16 || fieldValue.Kind() == reflect.Int32 ||
			fieldValue.Kind() == reflect.Int64 {
			if fieldValue.Int() == 0 {
				defaultValue := getDefaultValue(tagValue)
				// Convert the default value to the appropriate int type
				defaultIntValue, err := strconv.ParseInt(defaultValue, 10, 64)
				if err != nil {
					return fmt.Errorf("failed to convert default value to int: %w", err)
				}
				// Set the default value for the field
				fieldValue.SetInt(defaultIntValue)
			}
		} else if fieldValue.Kind() == reflect.Uint || fieldValue.Kind() == reflect.Uint8 ||
			fieldValue.Kind() == reflect.Uint16 || fieldValue.Kind() == reflect.Uint32 ||
			fieldValue.Kind() == reflect.Uint64 {
			if fieldValue.Uint() == 0 {
				defaultValue := getDefaultValue(tagValue)
				// Convert the default value to the appropriate uint type
				defaultUintValue, err := strconv.ParseUint(defaultValue, 10, 64)
				if err != nil {
					return fmt.Errorf("failed to convert default value to uint: %w", err)
				}
				// Set the default value for the field
				fieldValue.SetUint(defaultUintValue)
			}
		} else if fieldValue.Kind() == reflect.Float32 || fieldValue.Kind() == reflect.Float64 {
			if fieldValue.Float() == 0.0 {
				defaultValue := getDefaultValue(tagValue)
				// Convert the default value to the appropriate float type
				defaultFloatValue, err := strconv.ParseFloat(defaultValue, 64)
				if err != nil {
					return fmt.Errorf("failed to convert default value to float: %w", err)
				}
				// Set the default value for the field
				fieldValue.SetFloat(defaultFloatValue)
			}
		}
	}

	return nil
}

func getDefaultValue(tagValue string) string {
	// Workaround to using reflect again but for now it won't support spaces ;(
	defaultValue := strings.Split(tagValue, ":")[1]
	return strings.Split(defaultValue, " ")[0]
}
