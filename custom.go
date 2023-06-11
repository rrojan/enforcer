package enforcer

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type CustomEnforcements []map[string]func(string) string

func CustomValidator(req interface{}, customEnforcements CustomEnforcements) []string {
	errors := Validate(req)

	v := reflect.ValueOf(req)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		enforceTag := field.Tag.Get("enforce")

		if enforceTag != "" {
			fieldValue := v.Field(i)
			fieldString := ""
			fieldType := fieldValue.Type()
			switch fieldType.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				fieldString = strconv.Itoa(int(fieldValue.Int()))
			default:
				fieldString = fieldValue.String()
			}
			enforceOpts := strings.Split(enforceTag, " ")

			for _, opt := range enforceOpts {
				if strings.HasPrefix(opt, "custom") {
					enforcementNames := getCustomEnforcementNames(opt)
					for _, enforcementName := range enforcementNames {
						if enforcementFunc, ok := getCustomEnforcementFunc(customEnforcements, enforcementName); ok {
							// There's probably a better way to do all this
							err := enforcementFunc(fieldString)
							if err != "" {
								errors = append(errors, err)
							}
						} else {
							errors = append(errors, fmt.Sprintf("Custom enforcement '%s' not found for field '%s'", enforcementName, field.Name))
						}
					}
				}
			}
		}
	}

	return errors
}

func getCustomEnforcementFunc(
	customEnforcements CustomEnforcements, enforcementName string,
) (func(string) string, bool) {
	for _, enforcementMap := range customEnforcements {
		if enforcementFunc, ok := enforcementMap[enforcementName]; ok {
			return enforcementFunc, true
		}
	}
	return nil, false
}

func getCustomEnforcementNames(opt string) []string {
	prefix := "custom:"
	enforcementNamesStr := strings.TrimPrefix(opt, prefix)
	return strings.Split(enforcementNamesStr, ",")
}
