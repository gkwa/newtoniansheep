package core

import (
	"reflect"
	"testing"
)

func TestParseImageLink(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ImageLink
		wantErr  bool
	}{
		{
			name:     "Valid image link",
			input:    "![Alt text](https://example.com/image.jpg)",
			expected: ImageLink{Name: "Alt text", URL: "https://example.com/image.jpg"},
			wantErr:  false,
		},
		{
			name:     "Empty alt text",
			input:    "![](https://example.com/image.jpg)",
			expected: ImageLink{Name: "", URL: "https://example.com/image.jpg"},
			wantErr:  false,
		},
		{
			name:     "Invalid image link format",
			input:    "This is not an image link",
			expected: ImageLink{},
			wantErr:  false,
		},
		{
			name:     "Malformed image link",
			input:    "![Alt text](https://example.com/image.jpg",
			expected: ImageLink{},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseImageLink(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseImageLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ParseImageLink() = %v, want %v", got, tt.expected)
			}
		})
	}
}
