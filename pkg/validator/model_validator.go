package validator

import "github.com/go-playground/validator/v10"

var theValidator *validator.Validate

// GetValidator Flyweight for getting a structure validator.
func GetValidator() *validator.Validate {
	if theValidator == nil {
		theValidator = validator.New()
	}

	return theValidator
}
