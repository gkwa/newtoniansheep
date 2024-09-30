package core

import (
	"regexp"
	"strings"
)

type ImageLink struct {
	Name string
	URL  string
}

type ProcessorImpl struct {
	normalizer *LineNormalizer
}

func NewProcessor() *ProcessorImpl {
	return &ProcessorImpl{
		normalizer: NewLineNormalizer(),
	}
}

func (p *ProcessorImpl) Process(input []string) ([]string, error) {
	var result []string
	seenURLs := make(map[string]bool)

	for _, line := range input {
		trimmedLine := strings.TrimSpace(line)

		if strings.HasPrefix(trimmedLine, "[") || strings.HasPrefix(trimmedLine, "![") {
			imageLink, err := parseImageLink(trimmedLine)
			if err != nil {
				return nil, err
			}

			if !seenURLs[imageLink.URL] {
				seenURLs[imageLink.URL] = true
				result = append(result, line)
			}
		} else {
			result = append(result, line)
		}
	}

	return p.normalizer.Normalize(result), nil
}

func parseImageLink(line string) (ImageLink, error) {
	re := regexp.MustCompile(`\[([^\]]*)\]\(([^)]+)\)`)
	matches := re.FindStringSubmatch(line)
	if len(matches) == 3 {
		return ImageLink{
			Name: matches[1],
			URL:  matches[2],
		}, nil
	}
	return ImageLink{}, nil
}
