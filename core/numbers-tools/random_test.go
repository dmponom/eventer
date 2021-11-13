package numberstools

import (
	"testing"
)

func TestGetRandomIntInRange(t *testing.T) {
	tCases := []struct {
		name     string
		min, max int
		times    int
	}{
		{
			name:  "should generate random value for 10000 max",
			min:   100,
			max:   10000,
			times: 100,
		},
		{
			name:  "should generate random value for 1000000 max",
			min:   10000,
			max:   1000000,
			times: 100,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			prev := -1
			for i := 0; i < tc.times; i++ {
				result := GetRandomIntInRange(tc.min, tc.max)
				if result == prev {
					tt.Errorf("random repeated (result = %d, prev = %d)", result, prev)
				}

				if tc.min > result || result > tc.max {
					tt.Errorf("result is out of range (min = %d, max = %d, result = %d)", tc.min, tc.max, result)
				}

				prev = result
			}
		})
	}
}

func TestGetRandomInt(t *testing.T) {
	tCases := []struct {
		name  string
		maxN  int
		times int
	}{
		{
			name:  "should generate random value for 10000 max",
			maxN:  10000,
			times: 100,
		},
		{
			name:  "should generate random value for 1000000 max",
			maxN:  1000000,
			times: 100,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			prev := -1
			for i := 0; i < tc.times; i++ {
				result := GetRandomInt(tc.maxN)
				if result == prev {
					tt.Errorf("random repeated (result = %d, prev = %d)", result, prev)
				}

				if tc.maxN < result {
					tt.Errorf("result is greater than max (max = %d, result = %d)", tc.maxN, result)
				}

				prev = result
			}
		})
	}
}

func TestGetRandomIntSlice(t *testing.T) {
	tCases := []struct {
		name  string
		len   int
		times int
	}{
		{
			name:  "should generate slice of random value for len = 1000",
			len:   1000,
			times: 100,
		},
		{
			name:  "should generate slice of random value for len = 10000",
			len:   10000,
			times: 100,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			for i := 0; i < tc.times; i++ {
				result := GetRandomIntSlice(tc.len)
				if tc.len != len(result) {
					tt.Errorf("len(result) is not equal with expected (len(result) = %d, expected len = %d)", len(result), tc.len)
				}
			}
		})
	}
}
