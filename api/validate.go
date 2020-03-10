package api

import (
	"github.com/go-playground/validator/v10"
)

var v *validator.Validate

// Validate Validate
func Validate(i interface{}) error {
	v = validator.New()
	return v.Struct(i)
}
