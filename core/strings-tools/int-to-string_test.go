package stringtools

import (
	"testing"
)

func TestJoinInts(t *testing.T) {
	tCases := []struct {
		name     string
		nArr     []int
		expected string
	}{
		{
			name:     "should return empty string",
			nArr:     []int{},
			expected: "",
		},
		{
			name:     "should work with single value",
			nArr:     []int{1},
			expected: "1",
		},
		{
			name:     "should work with multi value",
			nArr:     []int{1, 2},
			expected: "1,2",
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := JoinInts(tc.nArr, ",")
			if actual != tc.expected {
				tt.Errorf("expected = %+v but actual = %+v", tc.expected, actual)
			}
		})
	}
}
