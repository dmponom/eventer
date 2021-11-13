package validators

import (
	numberTools "eventer/core/numbers-tools"
	stringTools "eventer/core/strings-tools"

	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
	"unicode"
)

func TestValidator(t *testing.T) {
	v, err := Make()
	if err != nil {
		t.Fatalf("can not build validator = %+v", err)
		return
	}

	if err = v.RegisterValidations(); err != nil {
		t.Fatalf("can not regiter validations = %+v", err)
		return
	}
	if err = v.RegisterTranslations(); err != nil {
		t.Fatalf("can not register translations = %+v", err)
		return
	}

	v.RegisterTagNameFunctions()

	type testStruct struct {
		TestEmail string `json:"test_email" validate:"required,min=8,max=25,EmailAddress"`
	}

	type testStructWithoutJSON struct {
		TestEmail string `validate:"required,min=8,max=25,EmailAddress"`
	}

	tCases := []struct {
		name           string
		entryStruct    testStruct
		expectedResult []string
	}{
		{
			name:           "Email field is valid",
			entryStruct:    testStruct{TestEmail: stringTools.GetRandomEmail()},
			expectedResult: nil,
		},
		{
			name:           "Email field is wrong",
			entryStruct:    testStruct{TestEmail: stringTools.GetRandom(12)},
			expectedResult: []string{"test_email: is wrong"},
		},
		{
			name:           "Email field is too short",
			entryStruct:    testStruct{TestEmail: "s@i.eu"},
			expectedResult: []string{"test_email: too short"},
		},
		{
			name:           "Email field is too long",
			entryStruct:    testStruct{TestEmail: fmt.Sprintf("%s%s", stringTools.GetRandom(25), stringTools.GetRandomEmail())},
			expectedResult: []string{"test_email: too long"},
		},
		{
			name:           "Email is not provided",
			entryStruct:    testStruct{TestEmail: ""},
			expectedResult: []string{"test_email: is required"},
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			var messages []string
			if err := v.GetValidator().Struct(tc.entryStruct); err != nil {
				valErrs := validator.ValidationErrors{}
				if errors.As(err, &valErrs) {
					for _, e := range valErrs {
						messages = append(messages, e.Translate(v.GetTranslator()))
					}
				}
			}
			if !cmp.Equal(messages, tc.expectedResult) {
				tt.Errorf("unexpected messages (actual = %+v expected = %+v)", messages, tc.expectedResult)
			}
		})
	}

	tCases2 := []struct {
		name           string
		entryStruct    testStructWithoutJSON
		expectedResult []string
	}{
		{
			name:           "testStructWithoutJSON: Email field is valid",
			entryStruct:    testStructWithoutJSON{TestEmail: stringTools.GetRandomEmail()},
			expectedResult: nil,
		},
		{
			name:           "testStructWithoutJSON: Email field is wrong",
			entryStruct:    testStructWithoutJSON{TestEmail: stringTools.GetRandom(12)},
			expectedResult: []string{"TestEmail: is wrong"},
		},
		{
			name:           "testStructWithoutJSON: Email field is too short",
			entryStruct:    testStructWithoutJSON{TestEmail: "s@i.eu"},
			expectedResult: []string{"TestEmail: too short"},
		},
		{
			name:           "testStructWithoutJSON: Email field is too long",
			entryStruct:    testStructWithoutJSON{TestEmail: fmt.Sprintf("%s%s", stringTools.GetRandom(25), stringTools.GetRandomEmail())},
			expectedResult: []string{"TestEmail: too long"},
		},
		{
			name:           "testStructWithoutJSON: Email is not provided",
			entryStruct:    testStructWithoutJSON{TestEmail: ""},
			expectedResult: []string{"TestEmail: is required"},
		},
	}

	for _, tc := range tCases2 {
		t.Run(tc.name, func(tt *testing.T) {
			var messages []string
			if err := v.GetValidator().Struct(tc.entryStruct); err != nil {
				valErrs := validator.ValidationErrors{}
				if errors.As(err, &valErrs) {
					for _, e := range valErrs {
						messages = append(messages, e.Translate(v.GetTranslator()))
					}
				}
			}
			if !cmp.Equal(messages, tc.expectedResult) {
				tt.Errorf("unexpected messages (actual = %+v expected = %+v)", messages, tc.expectedResult)
			}
		})
	}
}

func TestValidatorWithCustomFn(t *testing.T) {
	v, err := Make()
	if err != nil {
		t.Fatalf("can not build validator = %+v", err)
		return
	}

	if err = v.RegisterValidations(); err != nil {
		t.Fatalf("can not regiter validations = %+v", err)
		return
	}
	if err = v.RegisterTranslations(); err != nil {
		t.Fatalf("can not register translations = %+v", err)
		return
	}

	v.RegisterTagNameFunctions()

	type structWithCustomValidation struct {
		TestValue string `json:"test_value" validate:"OnlyDigit"`
	}

	customFunction := CustomFunction{
		Name:    "OnlyDigit",
		TextMsg: "string must have only digit symbols",
		Fn: func(value interface{}) bool {
			val := reflect.ValueOf(value)
			if val.Kind() != reflect.String {
				return false
			}

			for _, n := range val.String() {
				if !unicode.IsDigit(n) {
					return false
				}
			}
			return true
		},
	}

	if err = v.RegisterCustomValidation(customFunction); err != nil {
		t.Fatalf("can not register custom validation function = %+v", err)
		return
	}

	tCases := []struct {
		name           string
		entryStruct    structWithCustomValidation
		expectedResult []string
	}{
		{
			name:           "TestValue field is valid [have only digit symbols]",
			entryStruct:    structWithCustomValidation{TestValue: "123451234"},
			expectedResult: nil,
		},
		{
			name:           "TestValue field is not valid [have only chars symbols]",
			entryStruct:    structWithCustomValidation{TestValue: stringTools.GetRandom(12)},
			expectedResult: []string{"test_value: string must have only digit symbols"},
		},
		{
			name: "TestValue field is not valid [have mix digit & chars symbols]",
			entryStruct: structWithCustomValidation{
				TestValue: fmt.Sprintf("%s%d", stringTools.GetRandom(12), numberTools.GetRandomIntInRange(1, 10)),
			},
			expectedResult: []string{"test_value: string must have only digit symbols"},
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			var messages []string
			if err := v.GetValidator().Struct(tc.entryStruct); err != nil {
				valErrs := validator.ValidationErrors{}
				if errors.As(err, &valErrs) {
					for _, e := range valErrs {
						messages = append(messages, e.Translate(v.GetTranslator()))
					}
				}
			}
			if !cmp.Equal(messages, tc.expectedResult) {
				tt.Errorf("unexpected messages (actual = %+v expected = %+v)", messages, tc.expectedResult)
			}
		})
	}
}
