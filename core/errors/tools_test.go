package errors

import (
	"context"
	"errors"
	stringTools "eventer/core/strings-tools"
	"net/http"
	"testing"
)

func TestGetStatusCodeFromErr(t *testing.T) {
	testCases := []struct {
		name         string
		entryErr     error
		expectedCode int
	}{
		{
			name:         "GetStatusCodeFromErr: Success return code",
			entryErr:     InternalServerErrorBuilder(context.Background(), nil, stringTools.GetRandom(15)),
			expectedCode: http.StatusInternalServerError,
		},
		{
			name:         "GetStatusCodeFromErr: Failure return 0",
			entryErr:     errors.New(stringTools.GetRandom(15)),
			expectedCode: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			code := GetStatusCodeFromErr(tc.entryErr)
			if tc.expectedCode != code {
				tt.Fatalf("extected expectedCode = %d but actual expectedCode = %d", tc.expectedCode, code)
			}
		})
	}
}

func TestIsEqual(t *testing.T) {
	ctx := context.Background()
	err := errors.New(stringTools.GetRandom(15))

	errorMsg := stringTools.GetRandom(15)
	customError := InternalServerErrorBuilder(ctx, err, errorMsg)
	targetError := InternalServerErrorBuilder(ctx, err, errorMsg)
	otherError := CanNotUnmarshalErrorBuilder(ctx, err)

	testCases := []struct {
		name           string
		entryErr       error
		entryTarget    error
		expectedResult bool
	}{
		{
			name:           "IsEqual: should be return true (err == nil, target == nil)",
			entryErr:       nil,
			entryTarget:    nil,
			expectedResult: true,
		},
		{
			name:           "IsEqual: should be return false (err == nil, target == not nil)",
			entryErr:       nil,
			entryTarget:    targetError,
			expectedResult: false,
		},
		{
			name:           "IsEqual: should be return false (err == not nil, target == nil)",
			entryErr:       customError,
			entryTarget:    nil,
			expectedResult: false,
		},
		{
			name:           "IsEqual: should be return true",
			entryErr:       customError,
			entryTarget:    targetError,
			expectedResult: true,
		},

		{
			name:           "IsEqual: should be return false",
			entryErr:       customError,
			entryTarget:    otherError,
			expectedResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			if isEqual := IsEqual(tc.entryErr, tc.entryTarget); tc.expectedResult != isEqual {
				tt.Fatalf("extected expectedResult = %t but actual expectedResult = %t", tc.expectedResult, isEqual)
			}
		})
	}
}
