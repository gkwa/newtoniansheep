package core

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofrs/flock"
)

type FileHandlerImpl struct {
	lockTimeout time.Duration
}

func NewFileHandler() *FileHandlerImpl {
	return &FileHandlerImpl{
		lockTimeout: 10 * time.Second,
	}
}

func (fh *FileHandlerImpl) Read(path string) ([]string, error) {
	fileLock := flock.New(path + ".lock")
	ctx, cancel := context.WithTimeout(context.Background(), fh.lockTimeout)
	defer cancel()

	locked, err := fileLock.TryLockContext(ctx, 100*time.Millisecond)
	if err != nil {
		return nil, err
	}
	if !locked {
		return nil, context.DeadlineExceeded
	}
	defer func() {
		if err := fileLock.Unlock(); err != nil {
			fmt.Printf("Error unlocking file: %v\n", err)
		}
	}()

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(content), "\n"), nil
}

func (fh *FileHandlerImpl) Write(path string, content []string) error {
	fileLock := flock.New(path + ".lock")
	ctx, cancel := context.WithTimeout(context.Background(), fh.lockTimeout)
	defer cancel()

	locked, err := fileLock.TryLockContext(ctx, 100*time.Millisecond)
	if err != nil {
		return err
	}
	if !locked {
		return context.DeadlineExceeded
	}
	defer func() {
		if err := fileLock.Unlock(); err != nil {
			fmt.Printf("Error unlocking file: %v\n", err)
		}
	}()

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range content {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
