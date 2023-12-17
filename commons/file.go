package commons

import (
	"os"
	"path/filepath"
	"strings"
)

// EnsureDirExists checks if a directory exists at the given path, and creates it if it doesn't exist.
func EnsureDirExists(path string) error {

	path = CleanPath(path)
	path = Dir(path)

	if !PathExists(path) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// Extract Dir path from file path
func Dir(filePath string) string {
	return filepath.Dir(filePath)
}

// Path Exists
func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// CreateFile creates a file at the given path and writes the contents to the file.
// It overwrites the file if it already exists.
func CreateFile(path string, contents []byte) error {

	// Ensure the directory exists
	err := EnsureDirExists(path)
	if err != nil {
		return err
	}

	// Write contents to the file
	err = os.WriteFile(path, contents, 0644)
	if err != nil {
		return err
	}

	return nil
}

// PathPrepare replaces $HOME, $TMP, $CWD with their respective paths
func PathPrepare(path string) string {

	// replace $HOME with Home path if path contains $HOME is present
	if strings.HasPrefix(path, "$HOME") {
		path = strings.Replace(path, "$HOME", os.Getenv("HOME"), 1)
	}

	// replace $TMP with Temp path if path contains $TMP is present
	if strings.HasPrefix(path, "$TMP") {
		path = strings.Replace(path, "$TMP", os.Getenv("TMP"), 1)
	}

	// replace $CWD with current working path if path contains $CWD is present
	if strings.HasPrefix(path, "$CWD") {
		path = strings.Replace(path, "$CWD", os.Getenv("PWD"), 1)
	}

	// remove double slashes
	path = CleanPath(path)

	return path
}

// clean Path string
func CleanPath(path string) string {
	// remove double slashes
	path = strings.Replace(path, "//", "/", 1)

	return path
}

// Walk walks the file tree rooted at root, calling walkFn for each file or directory in the tree, including root.
func Walk(root string, walkFn filepath.WalkFunc) error {
	return filepath.Walk(root, walkFn)
}

// FileName from path
func FileName(path string) string {
	return filepath.Base(path)
}

// FilenameWithoutExtension from path
func FilenameWithoutExtension(path string) string {
	return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
}
