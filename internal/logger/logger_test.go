package logger

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

func TestNewConsoleLogger(t *testing.T) {
	tests := []struct {
		name           string
		development    bool
		enableStdout   bool
		expectedMsg    string
		expectedKey    string
		expectedValue  string
	}{
		{
			name:           "Development mode",
			development:    true,
			enableStdout:   true,
			expectedMsg:    "test message",
			expectedKey:    "key",
			expectedValue:  "value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldStderr := os.Stderr
			r, w, _ := os.Pipe()
			os.Stderr = w

			logger := NewConsoleLogger(tt.development, tt.enableStdout)
			logger.Info(tt.expectedMsg, tt.expectedKey, tt.expectedValue)

			w.Close()
			os.Stderr = oldStderr

			var buf bytes.Buffer
			_, err := buf.ReadFrom(r)
			if err != nil {
				t.Fatalf("Failed to read log output: %v", err)
			}

			var logMap map[string]interface{}
			err = json.Unmarshal(buf.Bytes(), &logMap)
			if err != nil {
				t.Fatalf("Failed to parse JSON log: %v", err)
			}

			if logMap["message"] != tt.expectedMsg {
				t.Errorf("Expected message '%s', got '%v'", tt.expectedMsg, logMap["message"])
			}

			if logMap[tt.expectedKey] != tt.expectedValue {
				t.Errorf("Expected key '%s' to be '%s', got '%v'", tt.expectedKey, tt.expectedValue, logMap[tt.expectedKey])
			}
		})
	}
}

