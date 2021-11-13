package validators

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	tCases := []struct {
		name           string
		emailEntry     string
		expectedResult bool
	}{
		{
			name:           "Regular email is valid",
			emailEntry:     "correct@email.com",
			expectedResult: true,
		},
		{
			name:           "Numerical email is valid",
			emailEntry:     "123@email.com",
			expectedResult: true,
		},
		{
			name:           "Email with subdomain is valid",
			emailEntry:     "123@subdomain.domain.com",
			expectedResult: true,
		},
		{
			name:           "Email with person name is valid",
			emailEntry:     "Some Person <still.correct@email.en>",
			expectedResult: true,
		},
		{
			name:           "Invalid email without '@' and second part",
			emailEntry:     "wrongemail",
			expectedResult: false,
		},
		{
			name:           "Invalid email without '@' and second part and with space in string",
			emailEntry:     "wrong email",
			expectedResult: false,
		},

		{
			name:           "Invalid email with wrong domain ending",
			emailEntry:     "worng@email.address",
			expectedResult: false,
		},
		{
			name:           "Invalid email without first part",
			emailEntry:     "@subdomain.domain.com",
			expectedResult: false,
		},
		{
			name:           "Invalid email without second part",
			emailEntry:     "123@",
			expectedResult: false,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			valid := IsEmailAddress(tc.emailEntry)

			if tc.expectedResult != valid {
				tt.Errorf("Method MakeConfig() return wrong config.token: got %v want %v", valid, tc.expectedResult)
			}
		})
	}
}
