package enforcements

import (
	"fmt"
	"reflect"
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
		if tagValue == "" {
			continue
		}

		// Check if the field is empty (zero value)
		if fieldValue.Kind() == reflect.String && fieldValue.String() == "" {
			// Workaround to using reflect again but for the meanwhile it 
			// wont support spaces ;(
			defaultValue := strings.Split(tagValue, ":")[1]
			defaultValue = strings.Split(defaultValue, " ")[0]
			// Set the default value for the field
			fieldValue.SetString(defaultValue)
		}
	}

	return nil
}