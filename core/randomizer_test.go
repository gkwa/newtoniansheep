package core

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/go-logr/logr"
)

func TestRandomizerImpl_Process(t *testing.T) {
	tests := []struct {
		name           string
		input          []string
		expectedImages int
		nonImageLines  map[int]string
	}{
		{
			name: "Basic randomization",
			input: []string{
				"![](http://example.com/image1.jpg)",
				"Some text 1",
				"![](http://example.com/image2.jpg)",
				"![](http://example.com/image3.jpg)",
				"Some text 2",
				"![](http://example.com/image4.jpg)",
				"![](http://example.com/image5.jpg)",
				"Some text 3",
				"![](http://example.com/image6.jpg)",
				"Some text 4",
			},
			expectedImages: 6,
			nonImageLines: map[int]string{
				1: "Some text 1",
				4: "Some text 2",
				7: "Some text 3",
				9: "Some text 4",
			},
		},
		{
			name: "No images",
			input: []string{
				"Some text 1",
				"Some text 2",
				"Some text 3",
			},
			expectedImages: 0,
			nonImageLines: map[int]string{
				0: "Some text 1",
				1: "Some text 2",
				2: "Some text 3",
			},
		},
		{
			name: "Only images",
			input: []string{
				"![](http://example.com/image1.jpg)",
				"![](http://example.com/image2.jpg)",
				"![](http://example.com/image3.jpg)",
			},
			expectedImages: 3,
			nonImageLines:  map[int]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRandomizer()
			var output []string
			var err error

			for i := 0; i < 5; i++ {
				output, err = r.Process(tt.input)
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}

				if !reflect.DeepEqual(tt.input, output) {
					break
				}

				time.Sleep(10 * time.Millisecond)
			}

			if reflect.DeepEqual(tt.input, output) && tt.expectedImages > 1 {
				t.Errorf(
					"Expected output to be different from input, but they are the same after 5 attempts",
				)
			}

			if len(output) != len(tt.input) {
				t.Errorf("Expected output length %d, got %d", len(tt.input), len(output))
			}

			imageLinks := 0
			for i, line := range output {
				if ImageLinkRegex.MatchString(line) {
					imageLinks++
				} else if expectedText, ok := tt.nonImageLines[i]; ok {
					if line != expectedText {
						t.Errorf("Non-image line at position %d has changed. Expected '%s', got '%s'", i, expectedText, line)
					}
				}
			}

			if imageLinks != tt.expectedImages {
				t.Errorf("Expected %d image links, got %d", tt.expectedImages, imageLinks)
			}
		})
	}
}

type mockRandomizerFileHandler struct {
	content []string
}

func (m *mockRandomizerFileHandler) Read(path string) ([]string, error) {
	return m.content, nil
}

func (m *mockRandomizerFileHandler) Write(path string, content []string) error {
	m.content = content
	return nil
}

func TestLinkRandomizer_ProcessFile(t *testing.T) {
	tests := []struct {
		name           string
		input          []string
		expectedImages int
		nonImageLines  map[int]string
	}{
		{
			name: "Basic file randomization",
			input: []string{
				"![](http://example.com/image1.jpg)",
				"Some text 1",
				"![](http://example.com/image2.jpg)",
				"![](http://example.com/image3.jpg)",
				"Some text 2",
			},
			expectedImages: 3,
			nonImageLines: map[int]string{
				1: "Some text 1",
				4: "Some text 2",
			},
		},
		{
			name: "File with no images",
			input: []string{
				"Some text 1",
				"Some text 2",
				"Some text 3",
			},
			expectedImages: 0,
			nonImageLines: map[int]string{
				0: "Some text 1",
				1: "Some text 2",
				2: "Some text 3",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHandler := &mockRandomizerFileHandler{content: tt.input}
			lr := NewLinkRandomizer(logr.Discard(), mockHandler, NewRandomizer())

			err := lr.ProcessFile("test.txt")
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			output := mockHandler.content

			if len(output) != len(tt.input) {
				t.Errorf("Expected output length %d, got %d", len(tt.input), len(output))
			}

			imageLinks := 0
			for i, line := range output {
				if ImageLinkRegex.MatchString(line) {
					imageLinks++
				} else if expectedText, ok := tt.nonImageLines[i]; ok {
					if line != expectedText {
						t.Errorf("Non-image line at position %d has changed. Expected '%s', got '%s'", i, expectedText, line)
					}
				}
			}

			if imageLinks != tt.expectedImages {
				t.Errorf("Expected %d image links, got %d", tt.expectedImages, imageLinks)
			}

			if reflect.DeepEqual(tt.input, output) && tt.expectedImages > 1 {
				t.Logf(
					"Warning: Output is the same as input for test case '%s'. This may occasionally happen due to randomization.",
					tt.name,
				)
			}
		})
	}
}

func TestLinkRandomizerE2E(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name: "Basic randomization",
			input: `![name](https://name.com)

![name1](https://name1.com)`,
		},
		{
			name: "Basic randomization with extra whitespace",
			input: `![name](https://name.com)


					![name1](https://name1.com)`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHandler := &mockRandomizerFileHandler{content: strings.Split(tt.input, "\n")}
			lr := NewLinkRandomizer(logr.Discard(), mockHandler, NewRandomizer())

			err := lr.ProcessFile("test.txt")
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			output := strings.Join(mockHandler.content, "\n")
			inputLines := strings.Split(tt.input, "\n")
			reversedInput := strings.Join(reverseSlice(inputLines), "\n")

			if output != tt.input && output != reversedInput {
				t.Errorf("Expected output to be either the same as input or reversed, but got:\n%s", output)
			}

			outputLines := strings.Split(output, "\n")

			if len(outputLines) != len(inputLines) {
				t.Errorf("Expected output to have %d lines, but got %d", len(inputLines), len(outputLines))
			}

			inputImages := make(map[string]bool)
			for _, line := range inputLines {
				if ImageLinkRegex.MatchString(line) {
					inputImages[line] = true
				}
			}

			for _, line := range outputLines {
				if ImageLinkRegex.MatchString(line) {
					if !inputImages[line] {
						t.Errorf("Unexpected image link in output: %s", line)
					}
					delete(inputImages, line)
				}
			}

			if len(inputImages) > 0 {
				t.Errorf("Some input images were not present in the output: %v", inputImages)
			}
		})
	}
}

func reverseSlice(s []string) []string {
	reversed := make([]string, len(s))
	for i, j := 0, len(s)-1; i < len(s); i, j = i+1, j-1 {
		reversed[i] = s[j]
	}
	return reversed
}
