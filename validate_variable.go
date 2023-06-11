package enforcer

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/rrojan/enforcer/enforcements"
)

// ValidateVar validates an individual variable based on the provided enforcement tag
func ValidateVar(value interface{}, enforceTag string) []string {
	v := reflect.ValueOf(value)
	t := v.Type()

	var errors []string

	// Convert the value to a string for validation
	var fieldValue string
	if v.Kind() == reflect.String {
		fieldValue = v.String()
	} else {
		fieldValue = fmt.Sprintf("%v", value)
	}

	enforceOpts := strings.Split(enforceTag, " ")

	for _, opt := range enforceOpts {
		switch {
		case opt == "required":
			err := enforcements.HandleRequired(fieldValue, "")
			if err != "" {
				errors = append(errors, err)
			}
		case strings.HasPrefix(opt, "between"):
			if t.Kind() == reflect.Int {
				err := enforcements.HandleBetweenInt(v.Int(), "", opt)
				if err != "" {
					errors = append(errors, err)
				}
			} else if t.Kind() == reflect.String {
				err := enforcements.HandleBetweenStr(fieldValue, "", opt)
				if err != "" {
					errors = append(errors, err)
				}
			} else {
				errors = append(errors, fmt.Sprintf("Unsupported type for between enforcement: %s", t.Kind()))
			}
		case strings.HasPrefix(opt, "min"):
			if t.Kind() == reflect.Int {
				err := enforcements.HandleMinInt(v.Int(), "", opt)
				if err != "" {
					errors = append(errors, err)
				}
			} else if t.Kind() == reflect.String {
				err := enforcements.HandleMinStr(fieldValue, "", opt)
				if err != "" {
					errors = append(errors, err)
				}
			} else {
				errors = append(errors, fmt.Sprintf("Unsupported type for min enforcement: %s", t.Kind()))
			}
		case strings.HasPrefix(opt, "max"):
			if t.Kind() == reflect.Int {
				err := enforcements.HandleMaxInt(v.Int(), "", opt)
				if err != "" {
					errors = append(errors, err)
				}
			} else if t.Kind() == reflect.String {
				err := enforcements.HandleMaxStr(fieldValue, "", opt)
				if err != "" {
					errors = append(errors, err)
				}
			} else {
				errors = append(errors, fmt.Sprintf("Unsupported type for max enforcement: %s", t.Kind()))
			}
		case strings.HasPrefix(opt, "match"):
			err := enforcements.HandleMatch(fieldValue, "", opt)
			if err != "" {
				errors = append(errors, err)
			}
		case strings.HasPrefix(opt, "wordCount"):
			err := enforcements.HandleWordCount(fieldValue, "", opt)
			if err != "" {
				errors = append(errors, err)
			}
		case strings.HasPrefix(opt, "enum"):
			if t.Kind() == reflect.Int {
				err := enforcements.HandleEnumIntOrFloat(v.Int(), "", opt)
				if err != "" {
					errors = append(errors, err)
				}
			} else if t.Kind() == reflect.Float32 || t.Kind() == reflect.Float64 {
				err := enforcements.HandleEnumIntOrFloat(v.Float(), "", opt)
				if err != "" {
					errors = append(errors, err)
				}
			} else if t.Kind() == reflect.String {
				err := enforcements.HandleEnumStr(fieldValue, "", opt)
				if err != "" {
					errors = append(errors, err)
				}
			} else {
				errors = append(errors, fmt.Sprintf("Unsupported type for enum enforcement: %s", t.Kind()))
			}
		case strings.HasPrefix(opt, "exclude"):
			if t.Kind() == reflect.Int {
				err := enforcements.HandleExcludeIntOrFloat(v.Int(), "", opt)
				if err != "" {
					errors = append(errors, err)
				}
			} else if t.Kind() == reflect.Float32 || t.Kind() == reflect.Float64 {
				err := enforcements.HandleExcludeIntOrFloat(v.Float(), "", opt)
				if err != "" {
					errors = append(errors, err)
				}
			} else if t.Kind() == reflect.String {
				err := enforcements.HandleExcludeStr(fieldValue, "", opt)
				if err != "" {
					errors = append(errors, err)
				}
			} else {
				errors = append(errors, fmt.Sprintf("Unsupported type for exclude enforcement: %s", t.Kind()))
			}
			// Add additional handlers for other enforcements as required
			// ...
		}
	}
	// Errors for single variable contains `''` because it is not a struct field
	for i := range errors {
		errors[i] = strings.ReplaceAll(errors[i], "''", "")
	}
	return errors
}