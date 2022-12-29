package fsutil

import (
	"os"
)

// Create a new File
func CreateFile(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return file, nil
}

// Check if a file exists
func PathExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}
