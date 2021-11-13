package stringtools

import (
	"github.com/google/uuid"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	testCases := []struct {
		name string
	}{
		{
			name: "should generate random uuid string",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			if _, err := uuid.Parse(generateUUID()); err != nil {
				tt.Errorf("expected error = <nil> but actual error = %v", err)
			}
		})
	}
}
