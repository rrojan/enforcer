package enforcer

import (
	"fmt"
	"reflect"
	"strings"
)

type CustomEnforcements []map[string]func(string) bool
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
			fieldValue := v.Field(i).String()
			enforceOpts := strings.Split(enforceTag, " ")

			for _, opt := range enforceOpts {
				switch {
				case strings.HasPrefix(opt, "custom"):
					enforcementNames := getCustomEnforcementNames(opt)
					for _, enforcementName := range enforcementNames {
						if enforcementFunc, ok := getCustomEnforcementFunc(customEnforcements, enforcementName); ok {
							isValid := enforcementFunc(fieldValue)
							if !isValid {
								errors = append(errors, fmt.Sprintf("Field '%s' failed custom validation '%s'", field.Name, enforcementName))
							}
						} else {
							errors = append(errors, fmt.Sprintf("Custom enforcement '%s' not found for field '%s'", enforcementName, field.Name))
						}
					}
				// Handle other enforcements
				// ...
				}
			}
		}
	}

	return errors
}

func getCustomEnforcementFunc(customEnforcements CustomEnforcements, enforcementName string) (func(string) bool, bool) {
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