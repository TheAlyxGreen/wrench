package path

import (
	"github.com/TheAlyxGreen/Wrench/data"
)

// NewPath returns a Path from an array of keys
func NewPath(keys []data.Data) Path {
	return Path{keys: keys}
}

// NewPathFromString returns a Path by parsing a string
func NewPathFromString(pathString string) Path {
	var currentKey string
	finalPath := Path{keys: []data.Data{}}
	escaped := false
	for _, c := range pathString {
		currentChar := string(c)
		if escaped {
			currentKey = currentKey + currentChar
			escaped = false
		} else {
			switch currentChar {
			case "/":
				PushKey(finalPath, data.NewDataFromString(currentKey))
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
		PushKey(finalPath, data.NewDataFromString(currentKey))
	}
	return finalPath
}
