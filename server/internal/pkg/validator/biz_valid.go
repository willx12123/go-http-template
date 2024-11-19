package validator

import (
	"github.com/go-playground/validator/v10"

	"server/internal/pkg/validator/validation"
)

func anyNamePattern(fl validator.FieldLevel) bool {
	return validation.ValidNameChar(fl.Field().String())
}
