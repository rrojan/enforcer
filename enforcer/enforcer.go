package enforcer

import (
	"reflect"
	"strings"

	"github.com/rrojan/enforcer/enforcer/enforcements"
)

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
			fieldValue := v.Field(i).String()
			enforceOpts := strings.Split(enforceTag, " ")

			for _, opt := range enforceOpts {
				switch {
				case opt == "required":
					err := enforcements.HandleRequired(fieldValue, field.Name)
					if err != "" {
						errors = append(errors, err)
					}
				case strings.HasPrefix(opt, "between"):
					err := enforcements.HandleBetween(fieldValue, field.Name, opt)
					if err != "" {
						errors = append(errors, err)
					}
				case strings.HasPrefix(opt, "match"):
					err := enforcements.HandleMatch(fieldValue, field.Name, opt)
					if err != "" {
						errors = append(errors, err)
					}
				// Add additional handlers for other enforcements as required
					// ...
				}
			}
		}
	}

	return errors
}
