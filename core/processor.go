package core

import (
	"strings"
)

type ProcessorImpl struct {
	normalizer *LineNormalizer
}

func NewProcessor() *ProcessorImpl {
	return &ProcessorImpl{
		normalizer: NewLineNormalizer(),
	}
}

func (p *ProcessorImpl) Process(input []string) ([]string, int, error) {
	var result []string
	seenURLs := make(map[string]bool)
	duplicatesRemoved := 0

	for _, line := range input {
		trimmedLine := strings.TrimSpace(line)

		if ImageLinkRegex.MatchString(trimmedLine) {
			imageLink, err := ParseImageLink(trimmedLine)
			if err != nil {
				return nil, 0, err
			}

			if !seenURLs[imageLink.URL] {
				seenURLs[imageLink.URL] = true
				result = append(result, line)
			} else {
				duplicatesRemoved++
			}
		} else {
			result = append(result, line)
		}
	}

	return p.normalizer.Normalize(result), duplicatesRemoved, nil
}
