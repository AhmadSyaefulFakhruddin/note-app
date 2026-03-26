package config

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var NotBlank validator.Func = func(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	return strings.TrimSpace(field) != ""
}
