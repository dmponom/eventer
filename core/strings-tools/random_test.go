package stringtools

import (
	"eventer/core/validators"
	"testing"

	numberTools "eventer/core/numbers-tools"
)

func TestGetRandom(t *testing.T) {
	tCases := []struct {
		name  string
		len   int
		times int
	}{
		{
			name:  "should generate random value for 30 max",
			len:   30,
			times: 100,
		},
		{
			name:  "should generate random value for 50 max",
			len:   50,
			times: 100,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			var prev string
			for i := 0; i < tc.times; i++ {
				result := GetRandom(tc.len)
				if result == prev {
					tt.Errorf("random repeated (result = %s, prev = %s)", result, prev)
				}

				if len(result) != tc.len {
					tt.Errorf("len(result) is not equal expected (expected = %d, len(result) = %d)", len(result), tc.len)
				}

				prev = result
			}
		})
	}
}

func TestGetRandomEmail(t *testing.T) {
	tCases := []struct {
		name    string
		isValid bool
		times   int
	}{
		{
			name:    "should return valid email",
			isValid: true,
			times:   100,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			for i := 0; i < tc.times; i++ {
				email := GetRandomEmail()
				isValid := validators.IsEmailAddress(email)
				if isValid != tc.isValid {
					tt.Errorf("expected isValid = %t but actual isValid = %t (email = %s)", tc.isValid, isValid, email)
				}
			}
		})
	}
}

func TestGetRandomEmails(t *testing.T) {
	tCases := []struct {
		name     string
		areValid bool
		times    int
		len      int
	}{
		{
			name:     "should return valid email",
			areValid: true,
			times:    100,
			len:      numberTools.GetRandomInt(100),
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			for i := 0; i < tc.times; i++ {
				emails := GetRandomEmails(tc.len)

				if len(emails) != tc.len {
					tt.Errorf("expected len = %d but actual len(emails) = %d", tc.len, len(emails))
				}

				for _, email := range emails {
					isValid := validators.IsEmailAddress(email)
					if isValid != tc.areValid {
						tt.Errorf("expected areValid = %t but actual isValid = %t (email = %s)", tc.areValid, isValid, email)
					}
				}
			}
		})
	}
}
