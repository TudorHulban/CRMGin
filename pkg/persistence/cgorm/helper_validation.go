package cgorm

import "github.com/TudorHulban/GinCRM/pkg/validator"

// validateStruct For unsigned integers use gte greater or equal for validation.
func validateStruct(data interface{}) error {
	return validator.GetValidator().Struct(data)
}
