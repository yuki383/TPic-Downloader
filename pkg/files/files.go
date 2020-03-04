// Package files provide routines of to create and write files.
package files

import (
	"errors"
	"os"
	"path/filepath"
)

// FileService has dependencies functions to execute file operations.
type FileService struct {
	joinPath func(elem ...string) string
	abs      func(path string) (string, error)
}

func newService() *FileService {
	s := &FileService{
		joinPath: filepath.Join,
		abs:      filepath.Abs,
	}

	return s
}

var defaultService = newService()

func New(directory string, name string, data []byte) (int, error) {
	if directory == "" || name == "" {
		err := errors.New("invalid arguments")
		return 0, err
	}

	return defaultService.new(directory, name, data)
}

func (f FileService) new(directory string, name string, data []byte) (int, error) {
	path, err := f.makeFilePath(directory, name)
	if err != nil {
		return 0, err
	}

	size, err := f.genarateFile(path, data)
	if err != nil {
		return 0, err
	}

	return size, nil

}

func (f FileService) makeFilePath(directory string, name string) (string, error) {
	p := f.joinPath(directory, name)
	abs, err := f.abs(p)
	if err != nil {
		return "", err
	}

	return abs, err
}

func (f FileService) genarateFile(path string, data []byte) (int, error) {
	file, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	size, err := file.Write(data)
	if err != nil {
		return 0, err
	}

	return size, nil
}
