package xmath

import (
	"testing"
)

func TestClip(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		minVal   any
		maxVal   any
		expected any
	}{
		// int tests
		{
			name:     "int value within range",
			value:    5,
			minVal:   1,
			maxVal:   10,
			expected: 5,
		},
		{
			name:     "int value below min",
			value:    0,
			minVal:   1,
			maxVal:   10,
			expected: 1,
		},
		{
			name:     "int value above max",
			value:    11,
			minVal:   1,
			maxVal:   10,
			expected: 10,
		},
		{
			name:     "int value at min boundary",
			value:    1,
			minVal:   1,
			maxVal:   10,
			expected: 1,
		},
		{
			name:     "int value at max boundary",
			value:    10,
			minVal:   1,
			maxVal:   10,
			expected: 10,
		},

		// float64 tests
		{
			name:     "float value within range",
			value:    3.5,
			minVal:   1.0,
			maxVal:   5.0,
			expected: 3.5,
		},
		{
			name:     "float value below min",
			value:    0.5,
			minVal:   1.0,
			maxVal:   5.0,
			expected: 1.0,
		},
		{
			name:     "float value above max",
			value:    6.0,
			minVal:   1.0,
			maxVal:   5.0,
			expected: 5.0,
		},
		{
			name:     "float value at min boundary",
			value:    1.0,
			minVal:   1.0,
			maxVal:   5.0,
			expected: 1.0,
		},
		{
			name:     "float value at max boundary",
			value:    5.0,
			minVal:   1.0,
			maxVal:   5.0,
			expected: 5.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got any
			switch v := tt.value.(type) {
			case int:
				got = Clip(v, tt.minVal.(int), tt.maxVal.(int))
			case float64:
				got = Clip(v, tt.minVal.(float64), tt.maxVal.(float64))
			}

			if got != tt.expected {
				t.Errorf("Clip() = %v, want %v", got, tt.expected)
			}
		})
	}
}
