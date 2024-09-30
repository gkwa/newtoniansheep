package core

import "github.com/go-logr/logr"

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

func (d *Deduplicator) ProcessFile(path string) error {
	content, err := d.fileHandler.Read(path)
	if err != nil {
		d.logger.Error(err, "Failed to read file")
		return err
	}

	processedContent, err := d.processor.Process(content)
	if err != nil {
		d.logger.Error(err, "Failed to process content")
		return err
	}

	err = d.fileHandler.Write(path, processedContent)
	if err != nil {
		d.logger.Error(err, "Failed to write file")
		return err
	}

	return nil
}
