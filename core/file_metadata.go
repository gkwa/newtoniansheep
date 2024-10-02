package core

import (
	"os"
	"regexp"

	"github.com/dustin/go-humanize"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type FileMetadata struct {
	Path              string
	Size              uint64
	LineCount         int
	LinkCount         int
	DuplicatesRemoved int
}

func GetFileMetadata(path string, duplicatesRemoved int) (FileMetadata, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return FileMetadata{}, err
	}

	lines := SplitLines(string(content))
	linkCount := CountLinks(lines)

	return FileMetadata{
		Path:              path,
		Size:              uint64(len(content)),
		LineCount:         len(lines),
		LinkCount:         linkCount,
		DuplicatesRemoved: duplicatesRemoved,
	}, nil
}

func SplitLines(s string) []string {
	return regexp.MustCompile(`\r?\n`).Split(s, -1)
}

func CountLinks(lines []string) int {
	linkRegex := regexp.MustCompile(`\[([^\]]*)\]\(([^)]+)\)`)
	count := 0
	for _, line := range lines {
		count += len(linkRegex.FindAllString(line, -1))
	}
	return count
}

func (fm FileMetadata) String() string {
	p := message.NewPrinter(language.English)
	size := humanize.Bytes(fm.Size)
	return p.Sprintf(
		"%s is %s with %d lines and %d links and %d duplicates removed",
		fm.Path,
		size,
		fm.LineCount,
		fm.LinkCount,
		fm.DuplicatesRemoved,
	)
}
