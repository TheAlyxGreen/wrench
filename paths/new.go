package paths

import (
	"github.com/TheAlyxGreen/wrench/values"
)

// New returns a Path by parsing a string
func New(pathString string) (Path, error) {
	path := RootPath
	if pathString == "" {
		return path, ErrorInvalidPath
	} else if pathString == rootPathString {
		return RootPath, nil
	}

	var currentKey string
	escaped := false
	for i, c := range pathString {
		currentChar := string(c)
		if escaped {
			currentKey = currentKey + currentChar
			escaped = false
		} else {
			switch currentChar {
			case pathDividerString:
				if currentKey != "" {
					path = Append(path, values.New(currentKey))
				} else if i != 0 {
					return Path{}, ErrorInvalidPath
				}
				currentKey = ""
				break
			case "\\":
				escaped = true
				break
			default:
				currentKey = currentKey + currentChar
			}
		}
	}
	if currentKey != "" {
		path = Append(path, values.New(currentKey))
	}
	return path, nil
}

// NewFromKeys returns a Path from an array of keys
func NewFromKeys(keys []values.Value) Path {
	return Path{keys: keys}
}
