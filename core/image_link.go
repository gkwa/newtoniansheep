package core

import "regexp"

var ImageLinkRegex = regexp.MustCompile(`!\[([^\]]*)\]\(([^)]+)\)`)

type ImageLink struct {
	Name string
	URL  string
}

func ParseImageLink(line string) (ImageLink, error) {
	matches := ImageLinkRegex.FindStringSubmatch(line)
	if len(matches) == 3 {
		return ImageLink{
			Name: matches[1],
			URL:  matches[2],
		}, nil
	}
	return ImageLink{}, nil
}
