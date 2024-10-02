package core

import (
	"reflect"
	"testing"
)

func TestLineNormalizer_Normalize(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Normal case",
			input:    []string{"line1", "", "", "line2", "  ", "line3", ""},
			expected: []string{"line1", "", "line2", "", "line3", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ln := NewLineNormalizer()
			result := ln.Normalize(tt.input)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
