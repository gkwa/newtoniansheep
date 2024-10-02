package core

import (
	"reflect"
	"testing"
)

func TestProcessorImpl_Process(t *testing.T) {
	tests := []struct {
		name               string
		input              []string
		expectedOutput     []string
		expectedDuplicates int
		expectedErr        bool
	}{
		{
			name: "Normal case",
			input: []string{
				"![](http://example.com/image1.jpg)",
				"![](http://example.com/image2.jpg)",
				"![](http://example.com/image1.jpg)",
				"Some text",
				"",
				"",
				"![](http://example.com/image3.jpg)",
			},
			expectedOutput: []string{
				"![](http://example.com/image1.jpg)",
				"![](http://example.com/image2.jpg)",
				"Some text",
				"",
				"![](http://example.com/image3.jpg)",
			},
			expectedDuplicates: 1,
			expectedErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewProcessor()
			output, duplicatesRemoved, err := p.Process(tt.input)

			if (err != nil) != tt.expectedErr {
				t.Fatalf("Unexpected error: %v", err)
			}

			if !reflect.DeepEqual(output, tt.expectedOutput) {
				t.Errorf("Expected %v, got %v", tt.expectedOutput, output)
			}

			if duplicatesRemoved != tt.expectedDuplicates {
				t.Errorf(
					"Expected %d duplicate removed, got %d",
					tt.expectedDuplicates,
					duplicatesRemoved,
				)
			}
		})
	}
}
