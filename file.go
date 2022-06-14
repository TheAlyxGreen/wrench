package wrench

import (
	"errors"
	"os"
)

// FilePath returns the location of the bolt File
func (w *Wrench) FilePath() string {
	return w.filePath
}

// fileExists verifies that a File exists
func fileExists(file string) bool {
	_, err := os.Stat(file)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

// FilePathExists verifies that the Wrench's filePath exists
func (w *Wrench) FilePathExists() bool {
	return fileExists(w.filePath)
}
