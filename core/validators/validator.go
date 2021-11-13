package validators

import (
	coreErrors "eventer/core/errors"
	"go.uber.org/fx"

	"context"
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/en"
	"reflect"
	"strings"
)

type driver struct {
	validator  *validator.Validate
	translator ut.Translator
}

func Make() (Validator, error) {
	ctx := context.Background()

	v := validator.New()
	t := en.New()
	uni := ut.New(t, t)

	trans, found := uni.GetTranslator(defaultLocale)
	if !found {
		return nil, coreErrors.InternalServerErrorBuilder(ctx, nil, fmt.Sprintf("default locale '%s' not found", defaultLocale))
	}

	if err := translations.RegisterDefaultTranslations(v, trans); err != nil {
		return nil, coreErrors.InternalServerErrorBuilder(ctx, err, "can not register default translations")
	}

	return &driver{
		validator:  v,
		translator: trans,
	}, nil
}

var Module = fx.Options(
	fx.Provide(Make),
	fx.Invoke(RegisterValidations, RegisterTranslations, RegisterTagNameFunctions),
)

func RegisterValidations(v Validator) error {
	return v.RegisterValidations()
}

func RegisterTranslations(v Validator) error {
	return v.RegisterTranslations()
}

func RegisterTagNameFunctions(v Validator) {
	v.RegisterTagNameFunctions()
}

func (d *driver) RegisterValidations() error {
	validators := []struct {
		name        string
		validatorFn func(fl validator.FieldLevel) bool
	}{
		{name: emailAddressFnName, validatorFn: emailAddress},
	}

	for _, val := range validators {
		if err := d.GetValidator().RegisterValidation(val.name, val.validatorFn); err != nil {
			return coreErrors.InternalServerErrorBuilder(context.Background(), err, "can not register validation")
		}
	}

	return nil
}

func (d *driver) RegisterTranslations() error {
	ctx := context.Background()

	translators := []struct {
		name string
		msg  string
	}{
		{name: requiredFnName, msg: "is required"},
		{name: minFnName, msg: "too short"},
		{name: maxFnName, msg: "too long"},
		{name: emailAddressFnName, msg: "is wrong"},
	}

	for _, translator := range translators {
		if err := d.registerTranslation(ctx, translator.name, translator.msg); err != nil {
			return err
		}
	}

	return nil
}

func (d *driver) RegisterCustomValidation(cf CustomFunction) error {
	ctx := context.Background()
	fn := func(fl validator.FieldLevel) bool {
		return cf.Fn(fl.Field().Interface())
	}

	// register custom function
	if err := d.GetValidator().RegisterValidation(cf.Name, fn); err != nil {
		return coreErrors.InternalServerErrorBuilder(ctx, err, "can not register custom validation")
	}

	// register translation for custom function
	return d.registerTranslation(ctx, cf.Name, cf.TextMsg)
}

func (d *driver) RegisterTagNameFunctions() {
	d.GetValidator().RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func (d *driver) GetValidator() *validator.Validate {
	return d.validator
}

func (d *driver) GetTranslator() ut.Translator {
	return d.translator
}

func (d *driver) registerTranslation(ctx context.Context, name, textMsg string) error {
	if err := d.GetValidator().RegisterTranslation(name, d.GetTranslator(), func(ut ut.Translator) error {
		return ut.Add(name, fmt.Sprintf("{0}: %s", textMsg), true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(name, fe.Field())
		return t
	}); err != nil {
		return coreErrors.InternalServerErrorBuilder(ctx, err, "can not register translation")
	}

	return nil
}

func emailAddress(fl validator.FieldLevel) bool {
	return IsEmailAddress(fl.Field().String())
}
