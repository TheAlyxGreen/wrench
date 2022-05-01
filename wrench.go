package wrench

import (
	"errors"
	"os"

	"github.com/TheAlyxGreen/wrench/path"
)

const ErrorFileDoesNotExist = "specified file does not exist"

// A Wrench is used to interact with a bolt database
type Wrench struct {
	filepath string
	cwd      path.Path
	ReadOnly bool
}

// FileExists verifies that a file exists
func FileExists(file string) bool {
	_, err := os.Stat(file)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

// NewWrench creates a new Wrench object from the specified file
func NewWrench(file string) (Wrench, error) {
	if !FileExists(file) {
		return Wrench{}, errors.New(ErrorFileDoesNotExist)
	}
	return Wrench{
		filepath: file,
		cwd:      path.NewPathFromString(""),
	}, nil
}

// GetFilepath returns the location of the bolt file
func (w *Wrench) GetFilepath() string {
	return w.filepath
}

// GetCWD returns the current position of the wrench in the bolt DB
func (w *Wrench) GetCWD() path.Path {
	return w.cwd
}

// FilepathExists verifies that the Wrench's filepath exists
func (w *Wrench) FilepathExists() bool {
	return FileExists(w.filepath)
}
