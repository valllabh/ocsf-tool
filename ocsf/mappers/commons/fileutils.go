package commons

import (
	"os"
	"path/filepath"
)

// EnsureDirExists checks if a directory exists at the given path, and creates it if it doesn't exist.
func EnsureDirExists(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateFile creates a file at the given path and writes the contents to the file.
// It overwrites the file if it already exists.
func CreateFile(filePath string, contents []byte) error {
	// Ensure the directory exists
	dir := filepath.Dir(filePath)
	err := EnsureDirExists(dir)
	if err != nil {
		return err
	}

	// Write contents to the file
	err = os.WriteFile(filePath, contents, 0644)
	if err != nil {
		return err
	}

	return nil
}
