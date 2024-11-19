package validator

import (
	"reflect"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type ginValidator struct {
	once     sync.Once
	validate *validator.Validate
}

func NewGinValidator() binding.StructValidator {
	return &ginValidator{
		validate: Default,
	}
}

func (v *ginValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}

func (v *ginValidator) Engine() interface{} {
	return v.validate
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
