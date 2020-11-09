package zeta

import "github.com/go-playground/validator/v10"

type Validator interface {
	Validate(s interface{}) error
}

type defaultValidator struct {
	validator *validator.Validate
}

func (v defaultValidator) Validate(s interface{}) error{
	return v.validator.Struct(s)
}

func DefaultValidator() Validator{
	return defaultValidator{validator.New()}
}