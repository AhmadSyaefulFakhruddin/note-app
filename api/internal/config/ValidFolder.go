package config

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var ValidFolder validator.Func = func(fl validator.FieldLevel) bool {
	field := fl.Field().String()

	if field == "" {
		return false
	}

	if !strings.HasPrefix(field, "/") || strings.HasSuffix(field, "/") {
		return false
	}

	if strings.Contains(field, "//") {
		return false
	}

	return true
}
