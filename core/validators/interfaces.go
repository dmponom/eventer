package validators

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Validator interface {
	RegisterValidations() error
	RegisterTranslations() error
	RegisterTagNameFunctions()

	RegisterCustomValidation(cv CustomFunction) error

	GetValidator() *validator.Validate
	GetTranslator() ut.Translator
}
