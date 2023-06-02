package enforcer

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Validate validates the fields of the given struct based on the `enforce` tags
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

		if enforceTag == "" {
			return []string{}
		}

		fieldValue := v.Field(i).String()
		enforceOpts := strings.Split(enforceTag, " ")

		for _, opt := range enforceOpts {
			if opt == "required" && fieldValue == "" {
				errors = append(errors, "Required field "+field.Name+" is absent")
			}

			if strings.HasPrefix(opt, "between") {
				rangeVals := strings.Split(strings.TrimPrefix(opt, "between:"), ",")
				fmt.Println("---")
				fmt.Printf("\n%+v\n", rangeVals)
				fmt.Printf("\n%+v\n", fieldValue)
				fmt.Println("---")
				
				min, err1 := strconv.Atoi(rangeVals[0])
				max, err2 := strconv.Atoi(rangeVals[1])
				
				if len(rangeVals) != 2 || err1 != nil || err2 != nil {
					errors = append(errors, "Invalid range values for field "+field.Name)
					continue
				}
				if len(fieldValue) < min || len(fieldValue) > max {
					errors = append(errors, "Field "+field.Name+" must be between "+rangeVals[0]+" and "+rangeVals[1]+" characters")
				}
			}

			if strings.HasPrefix(opt, "match") {
				pattern := strings.TrimPrefix(opt, "match:")
				match, err := regexp.MatchString(pattern, fieldValue)
				if err != nil {
					errors = append(errors, "Invalid pattern for field "+field.Name)
					continue
				}

				if !match {
					errors = append(errors, "Field "+field.Name+" does not match pattern")
				}
			}
		}
		
	}

	return errors
}
