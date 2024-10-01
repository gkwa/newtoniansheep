package core

import (
	"fmt"
	"path/filepath"

	"github.com/go-logr/logr"
)

type DeduplicateManager struct {
	logger      logr.Logger
	fileHandler FileHandler
	processor   Processor
}

func NewDeduplicateManager(logger logr.Logger, fileHandler FileHandler, processor Processor) *DeduplicateManager {
	return &DeduplicateManager{
		logger:      logger,
		fileHandler: fileHandler,
		processor:   processor,
	}
}

func (dm *DeduplicateManager) Deduplicate(inputPath string) (string, error) {
	absInputPath, err := filepath.Abs(inputPath)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute input file path: %w", err)
	}

	deduplicator := NewDeduplicator(dm.logger, dm.fileHandler, dm.processor)
	duplicatesRemoved, err := deduplicator.ProcessFile(absInputPath)
	if err != nil {
		return "", fmt.Errorf("failed to process file: %w", err)
	}

	metadata, err := GetFileMetadata(absInputPath, duplicatesRemoved)
	if err != nil {
		return "", fmt.Errorf("failed to get file metadata: %w", err)
	}

	return metadata.String(), nil
}
