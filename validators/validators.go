package validators

import (
	"reflect"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

// ValidateInputs -> sanitize user inputs
func ValidateInputs(dataSet interface{}) (bool, map[string][]string) {
	validate = validator.New()
	err := validate.Struct(dataSet)

	if err != nil {

		//Validation syntax is invalid
		if err, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}

		//Validation errors occurred
		errors := make(map[string][]string)
		//Use reflector to reverse engineer struct
		reflected := reflect.ValueOf(dataSet)
		for _, err := range err.(validator.ValidationErrors) {

			// Attempt to find field by name and get json tag name
			field, _ := reflected.Type().FieldByName(err.StructField())
			var name string

			//If json tag doesn't exist, use lower case of name
			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}

			switch err.Tag() {
			case "required":
				errors[name] = append(errors[name], "The "+name+" is required")
				break
			case "alpha":
				errors[name] = append(errors[name], "The "+name+" should contain only letters")
				break
			// case "eqfield":
			// 	errors[name] = append(errors[name], "The "+name+" should be equal to the "+err.Param())
			// 	break
			default:
				errors[name] = append(errors[name], "The "+name+" is invalid")
				break
			}
		}

		return false, errors
	}
	return true, nil
}
