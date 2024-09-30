package core

import "strings"

type LineNormalizer struct{}

func NewLineNormalizer() *LineNormalizer {
	return &LineNormalizer{}
}

func (ln *LineNormalizer) Normalize(lines []string) []string {
	var result []string
	var consecutiveEmptyLines int

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			consecutiveEmptyLines++
			if consecutiveEmptyLines <= 1 {
				result = append(result, "")
			}
		} else {
			consecutiveEmptyLines = 0
			result = append(result, line)
		}
	}

	return result
}
