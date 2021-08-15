package service

import (
	"io/ioutil"
)

// FileService type for creating files.
type FileService struct {
}

// writeJSONFile writes a value in key.json file.
func writeJSONFile(key string, value string) error {
	filename := key + ".json"
	valueBytes := []byte(value)

	err := ioutil.WriteFile(filename, valueBytes, 0644)
	return err
}

// Create creates a json file.
func (file *FileService) Create(key string, value string) error {
	err := writeJSONFile(key, value)
	return err
}

// Read reads a json file.
func (file *FileService) Read(key string) (string, error) {
	filename := key + ".json"

	valueBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(valueBytes), nil
}

// Update updates a json file.
func (file *FileService) Update(key string, value string) error {
	err := writeJSONFile(key, value)
	return err
}
