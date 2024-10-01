package core

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/go-logr/logr"
)

type Processor interface {
	Process(input []string) ([]string, int, error)
}

type FileHandler interface {
	Read(path string) ([]string, error)
	Write(path string, content []string) error
}

type Deduplicator struct {
	logger      logr.Logger
	fileHandler FileHandler
	processor   Processor
}

func NewDeduplicator(
	logger logr.Logger,
	fileHandler FileHandler,
	processor Processor,
) *Deduplicator {
	return &Deduplicator{
		logger:      logger,
		fileHandler: fileHandler,
		processor:   processor,
	}
}

func (d *Deduplicator) ProcessFile(path string) (int, error) {
	content, err := d.fileHandler.Read(path)
	if err != nil {
		d.logger.Error(err, "Failed to read file")
		return 0, err
	}

	originalChecksum := d.calculateChecksum(content)

	processedContent, duplicatesRemoved, err := d.processor.Process(content)
	if err != nil {
		d.logger.Error(err, "Failed to process content")
		return 0, err
	}

	processedChecksum := d.calculateChecksum(processedContent)

	if originalChecksum == processedChecksum {
		d.logger.V(1).Info("No changes detected, skipping file write")
		return duplicatesRemoved, nil
	}

	err = d.fileHandler.Write(path, processedContent)
	if err != nil {
		d.logger.Error(err, "Failed to write file")
		return 0, err
	}

	return duplicatesRemoved, nil
}

func (d *Deduplicator) calculateChecksum(content []string) string {
	hash := sha256.New()
	hash.Write([]byte(strings.Join(content, "\n")))
	return hex.EncodeToString(hash.Sum(nil))
}
