package core

import (
	"testing"

	"github.com/go-logr/logr"
)

type mockFileHandler struct{}

func (m *mockFileHandler) Read(path string) ([]string, error) {
	return []string{"line1", "line2", "line2", "line3"}, nil
}

func (m *mockFileHandler) Write(path string, content []string) error {
	return nil
}

type mockProcessor struct{}

func (m *mockProcessor) Process(input []string) ([]string, int, error) {
	return []string{"line1", "line2", "line3"}, 1, nil
}

func TestDeduplicator_ProcessFile(t *testing.T) {
	tests := []struct {
		name               string
		expectedDuplicates int
		expectedError      bool
	}{
		{
			name:               "Normal case",
			expectedDuplicates: 1,
			expectedError:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDeduplicator(logr.Discard(), &mockFileHandler{}, &mockProcessor{})
			duplicatesRemoved, err := d.ProcessFile("test.txt")

			if (err != nil) != tt.expectedError {
				t.Fatalf("Unexpected error status: %v", err)
			}
			if duplicatesRemoved != tt.expectedDuplicates {
				t.Errorf(
					"Expected %d duplicates removed, got %d",
					tt.expectedDuplicates,
					duplicatesRemoved,
				)
			}
		})
	}
}
