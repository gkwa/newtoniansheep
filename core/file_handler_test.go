package core

import (
	"os"
	"reflect"
	"testing"
)

func TestFileHandlerImpl_ReadWrite(t *testing.T) {
	tests := []struct {
		name        string
		content     []string
		expectedErr bool
	}{
		{
			name:        "Normal case",
			content:     []string{"line1", "line2", "line3"},
			expectedErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fh := NewFileHandler()
			testFile := "test_file.txt"

			err := fh.Write(testFile, tt.content)
			if (err != nil) != tt.expectedErr {
				t.Fatalf("Write error: %v", err)
			}

			readContent, err := fh.Read(testFile)
			if (err != nil) != tt.expectedErr {
				t.Fatalf("Read error: %v", err)
			}

			if len(readContent) > 0 && readContent[len(readContent)-1] == "" {
				readContent = readContent[:len(readContent)-1]
			}

			if !reflect.DeepEqual(readContent, tt.content) {
				t.Errorf("Expected %v, got %v", tt.content, readContent)
			}

			os.Remove(testFile)
		})
	}
}

