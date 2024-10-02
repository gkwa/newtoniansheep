package core

import (
	"os"
	"testing"
)

func TestGetFileMetadata(t *testing.T) {
	tests := []struct {
		name            string
		content         string
		expectedMinSize uint64
		expectedMaxSize uint64
		expectedLine    int
		expectedLink    int
		expectedErr     bool
	}{
		{
			name:            "Normal case",
			content:         "line1\nline2\n[link](http://example.com)",
			expectedMinSize: 35, // Minimum expected size (Unix line endings)
			expectedMaxSize: 38, // Maximum expected size (Windows line endings)
			expectedLine:    3,
			expectedLink:    1,
			expectedErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFile := "test_metadata.txt"

			err := os.WriteFile(testFile, []byte(tt.content), 0o644)
			if err != nil {
				t.Fatalf("Failed to write test file: %v", err)
			}
			defer os.Remove(testFile)

			metadata, err := GetFileMetadata(testFile, 0)
			if (err != nil) != tt.expectedErr {
				t.Fatalf("Unexpected error: %v", err)
			}

			if metadata.Path != testFile {
				t.Errorf("Expected path %s, got %s", testFile, metadata.Path)
			}
			if metadata.Size < tt.expectedMinSize || metadata.Size > tt.expectedMaxSize {
				t.Errorf(
					"Expected size between %d and %d, got %d",
					tt.expectedMinSize,
					tt.expectedMaxSize,
					metadata.Size,
				)
			}
			if metadata.LineCount != tt.expectedLine {
				t.Errorf("Expected %d lines, got %d", tt.expectedLine, metadata.LineCount)
			}
			if metadata.LinkCount != tt.expectedLink {
				t.Errorf("Expected %d link, got %d", tt.expectedLink, metadata.LinkCount)
			}
		})
	}
}
