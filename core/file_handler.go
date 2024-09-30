package core

import (
	"bufio"
	"os"
	"strings"
	"sync"
)

type FileHandlerImpl struct {
	mutex sync.Mutex
}

func NewFileHandler() *FileHandlerImpl {
	return &FileHandlerImpl{}
}

func (fh *FileHandlerImpl) Read(path string) ([]string, error) {
	fh.mutex.Lock()
	defer fh.mutex.Unlock()

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(content), "\n"), nil
}

func (fh *FileHandlerImpl) Write(path string, content []string) error {
	fh.mutex.Lock()
	defer fh.mutex.Unlock()

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
