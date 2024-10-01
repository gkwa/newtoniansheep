package core

import (
	"fmt"
	"path/filepath"

	"github.com/go-logr/logr"
)

type RandomizeManager struct {
	logger      logr.Logger
	fileHandler FileHandler
	processor   Randomizer
}

func NewRandomizeManager(logger logr.Logger, fileHandler FileHandler, processor Randomizer) *RandomizeManager {
	return &RandomizeManager{
		logger:      logger,
		fileHandler: fileHandler,
		processor:   processor,
	}
}

func (rm *RandomizeManager) Randomize(inputPath string) (string, error) {
	absInputPath, err := filepath.Abs(inputPath)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute input file path: %w", err)
	}

	outputPath := GetRandomizedFilePath(absInputPath)
	absOutputPath, err := filepath.Abs(outputPath)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute output file path: %w", err)
	}

	randomizer := NewLinkRandomizer(rm.logger, rm.fileHandler, rm.processor)
	err = randomizer.ProcessFile(absInputPath)
	if err != nil {
		return "", fmt.Errorf("failed to process file: %w", err)
	}

	metadata, err := GetFileMetadata(absOutputPath, 0)
	if err != nil {
		return "", fmt.Errorf("failed to get file metadata: %w", err)
	}

	return metadata.String(), nil
}
