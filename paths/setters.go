package paths

import (
	"github.com/TheAlyxGreen/wrench/values"
)

// Append appends the given Value to the end of the paths
func Append(p Path, key values.Value) Path {
	p.keys = append(p.keys, key)
	return p
}

// AppendString creates a new key from the string and appends it to the path
func AppendString(p Path, keyName string) Path {
	p.keys = append(p.keys, values.New(keyName))
	return p
}

// AppendPath appends all the keys from one path onto another paths
func AppendPath(p Path, p2 Path) Path {
	for _, key := range p2.keys {
		p.keys = append(p.keys, key)
	}
	return p
}

// AppendKeys appends an array of keys onto the paths
func AppendKeys(p Path, keys []values.Value) Path {
	for _, key := range keys {
		p.keys = append(p.keys, key)
	}
	return p
}
