package validator

import (
	"log"

	"github.com/go-playground/validator/v10"
)

var Default = validator.New()

func Init() {
	registerValidation("anynamepattern", anyNamePattern)

	Default.RegisterAlias("anyname", "required,max=16,min=1,anynamepattern")
}

func registerValidation(tag string, fn validator.Func) {
	err := Default.RegisterValidation(tag, fn)
	if err != nil {
		log.Panicf("register validation fail: %s", err)
	}
}
