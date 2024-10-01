package core

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"strings"

	"github.com/go-logr/logr"
)

type Randomizer interface {
	Process(input []string) ([]string, error)
}

type LinkRandomizer struct {
	logger      logr.Logger
	fileHandler FileHandler
	processor   Randomizer
}

func NewLinkRandomizer(
	logger logr.Logger,
	fileHandler FileHandler,
	processor Randomizer,
) *LinkRandomizer {
	return &LinkRandomizer{
		logger:      logger,
		fileHandler: fileHandler,
		processor:   processor,
	}
}

func (r *LinkRandomizer) ProcessFile(path string) error {
	content, err := r.fileHandler.Read(path)
	if err != nil {
		r.logger.Error(err, "Failed to read file")
		return err
	}

	processedContent, err := r.processor.Process(content)
	if err != nil {
		r.logger.Error(err, "Failed to process content")
		return err
	}

	outputPath := GetRandomizedFilePath(path)
	err = r.fileHandler.Write(outputPath, processedContent)
	if err != nil {
		r.logger.Error(err, "Failed to write file")
		return err
	}

	return nil
}

func GetRandomizedFilePath(originalPath string) string {
	dir := filepath.Dir(originalPath)
	filename := filepath.Base(originalPath)
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)
	return filepath.Join(dir, fmt.Sprintf("%s-randomized%s", name, ext))
}

type RandomizerImpl struct {
	rng *rand.Rand
}

func NewRandomizer() *RandomizerImpl {
	return &RandomizerImpl{
		rng: rand.New(rand.NewSource(rand.Int63())),
	}
}

func (r *RandomizerImpl) Process(input []string) ([]string, error) {
	var result []string
	var imageLinks []string
	var imageLinkIndices []int

	for i, line := range input {
		if ImageLinkRegex.MatchString(line) {
			imageLinks = append(imageLinks, line)
			imageLinkIndices = append(imageLinkIndices, i)
		}
	}

	r.rng.Shuffle(len(imageLinks), func(i, j int) {
		imageLinks[i], imageLinks[j] = imageLinks[j], imageLinks[i]
	})

	result = make([]string, len(input))
	copy(result, input)

	for i, index := range imageLinkIndices {
		result[index] = imageLinks[i]
	}

	return result, nil
}
