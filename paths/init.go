package paths

import (
	"errors"

	"github.com/TheAlyxGreen/wrench/values"
)

const rootPathString = "/"
const pathDividerString = "/"

var RootPath Path

var ErrorInvalidPath error
var ErrorPathHasNoKeys error

func init() {
	ErrorInvalidPath = errors.New("specified paths string is not valid")
	ErrorPathHasNoKeys = errors.New("paths has no keys")
	RootPath = Path{keys: []values.Value{}}
}
