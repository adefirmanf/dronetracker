package handler

import (
	"errors"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func (v *Validator) Validate(i interface{}) error {
	var errs []error
	if err := v.validator.Struct(i); err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			for _, err := range err.(validator.ValidationErrors) {
				errs = append(errs, errorTranslation(err))
			}
		}
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func errorTranslation(err validator.FieldError) error {
	switch err.Tag() {
	case "gte":
		return fmt.Errorf("%v must be greater than or equal to %v ", err.Field(), err.Param())
	case "lte":
		return fmt.Errorf("%v must be less than or equal to %v ", err.Field(), err.Param())
	case "uuid":
		return fmt.Errorf("%v must be a valid UUID", err.Field())
	}
	log.Println(err.Tag())
	return err
}
