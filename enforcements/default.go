package enforcements

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
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

		if strings.Contains(tagValue, "prohibit") {
			// If we are using prohibit with this field, reset the value
			// to whatever the Zero value of that type is as a default
			fieldValue.Set(reflect.Zero(fieldType.Type))
		}

		if tagValue == "" || !strings.Contains(tagValue, "default:") {
			continue
		}

		// Check if the field is empty (zero value)
		switch fieldValue.Kind() {
		case reflect.String:
			if fieldValue.String() == "" {
				defaultValue := getDefaultValue(tagValue)
				// Set the default value for the field
				fieldValue.SetString(defaultValue)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
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
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
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
		case reflect.Float32, reflect.Float64:
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
		case reflect.Struct:
			if fieldValue.Type() == reflect.TypeOf(time.Time{}) {
				if fieldValue.Interface().(time.Time).IsZero() {
					defaultValue := getDefaultValue(tagValue)
					defaultValue = strings.ReplaceAll(defaultValue, ";", ":")
					// Parse the default value as a time string
					defaultTime, err := time.Parse("2006-01-02 15:04:05 -07:00", defaultValue)
					if err != nil {
						return fmt.Errorf("failed to convert default value to time: %w", err)
					}

					// Set the default value for the field
					fieldValue.Set(reflect.ValueOf(defaultTime))
				}
			}
		}
	}

	return nil
}

func getDefaultValue(tagValue string) string {
	defaultValue := ""
	// if strings.Contains(tagValue, "'") {
	// 	re := regexp.MustCompile(`'([^']*)'`)
	// 	match := re.FindStringSubmatch(tagValue)

	// 	if len(match) >= 2 {
	// 		result := match[1]
	// 		fmt.Println(result)
	// 	}
	// }
	defaultValue = strings.Split(tagValue, ":")[1]
	return strings.TrimSpace(defaultValue)
	// return strings.Split(defaultValue, " ")[0]
	// TODO: this can mess things up
}