package echo

import (
	validator "github.com/go-playground/validator/v10"
)

// CutstomValidator :
type customValidator struct {
	Validator *validator.Validate
}

// Validate : Validate Data
func (cv *customValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
