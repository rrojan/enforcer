package enforcer

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/rrojan/enforcer/enforcer/enforcements"
)

func CustomValidator(req interface{}, customEnforcements []map[string]func(string) bool) []string {
	return []string{}
}

// Validate fields of a given struct based on `enforce` tags
func Validate(req interface{}) []string {
	v := reflect.ValueOf(req)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	var errors []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		enforceTag := field.Tag.Get("enforce")

		if enforceTag != "" {
			fieldValue := v.Field(i)
			fieldType := fieldValue.Type()
			fieldString := fieldValue.String()
			enforceOpts := strings.Split(enforceTag, " ")

			for _, opt := range enforceOpts {
				switch {
				case opt == "required":
					err := enforcements.HandleRequired(fieldString, field.Name)
					if err != "" {
						errors = append(errors, err)
					}
				case strings.HasPrefix(opt, "between"):
					if fieldType.Kind() == reflect.Int {
						err := enforcements.HandleBetweenInt(fieldValue.Int(), field.Name, opt)
						if err != "" {
							errors = append(errors, err)
						}
					} else if fieldType.Kind() == reflect.String {
						err := enforcements.HandleBetweenStr(fieldString, field.Name, opt)
						if err != "" {
							errors = append(errors, err)
						}
					} else {
						errors = append(errors, fmt.Sprintf("Unsupported type for field '%s'", field.Name))
					}
				case strings.HasPrefix(opt, "min"):
					if fieldType.Kind() == reflect.Int {
						err := enforcements.HandleMinInt(fieldValue.Int(), field.Name, opt)
						if err != "" {
							errors = append(errors, err)
						}
					} else if fieldType.Kind() == reflect.String {
						err := enforcements.HandleMinStr(fieldString, field.Name, opt)
						if err != "" {
							errors = append(errors, err)
						}
					} else {
						errors = append(errors, fmt.Sprintf("Unsupported type for field '%s'", field.Name))
					}
				case strings.HasPrefix(opt, "max"):
					if fieldType.Kind() == reflect.Int {
						err := enforcements.HandleMaxInt(fieldValue.Int(), field.Name, opt)
						if err != "" {
							errors = append(errors, err)
						}
					} else if fieldType.Kind() == reflect.String {
						err := enforcements.HandleMaxStr(fieldString, field.Name, opt)
						if err != "" {
							errors = append(errors, err)
						}
					} else {
						errors = append(errors, fmt.Sprintf("Unsupported type for field '%s'", field.Name))
					}
				case strings.HasPrefix(opt, "match"):
					err := enforcements.HandleMatch(fieldString, field.Name, opt)
					if err != "" {
						errors = append(errors, err)
					}
				case strings.HasPrefix(opt, "enum"):
					if fieldType.Kind() == reflect.Int {
						err := enforcements.HandleEnumIntOrFloat(fieldValue.Int(), field.Name, opt)
						if err != "" {
							errors = append(errors, err)
						}
					} else if fieldType.Kind() == reflect.Float32 || fieldType.Kind() == reflect.Float64 {
						err := enforcements.HandleEnumIntOrFloat(fieldValue.Float(), field.Name, opt)
						if err != "" {
							errors = append(errors, err)
						}
					} else if fieldType.Kind() == reflect.String {
						err := enforcements.HandleEnumStr(fieldString, field.Name, opt)
						if err != "" {
							errors = append(errors, err)
						}
					} else {
						errors = append(errors, fmt.Sprintf("Unsupported type for field '%s'", field.Name))
					}
					// case strings.HasPrefix(opt, "exclude"):
					// 	err := enforcements.HandleExclude(fieldString, field.Name, opt)
					// 	if err != "" {
					// 		errors = append(errors, err)
					// 	}
					// Add additional handlers for other enforcements as required
					// ...
				}
			}
		}
	}

	return errors
}
