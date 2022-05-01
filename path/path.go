package path

import (
	"github.com/TheAlyxGreen/wrench/data"
)

const ErrorPathHasNoKeys = "path has no keys"

// A Path is the series of buckets that must be traversed to arrive at a specific key
type Path struct {
	keys []data.Data
}

// ToString returns the path as a parseable path string
func ToString(p Path) string {
	if len(p.keys) == 0 {
		return "/"
	}
	var pathString string
	for _, k := range p.keys {
		pathString = "/" + pathString + data.ToKey(k)
	}
	return pathString
}

// Equals returns true if both paths contain the same keys in the same order
func Equals(p1, p2 Path) bool {
	if len(p1.keys) != len(p2.keys) {
		return false
	} else if len(p1.keys) == 0 {
		return true
	}
	for i, key := range p1.keys {
		if !data.Equals(key, p2.keys[i]) {
			return false
		}
	}
	return true
}

// Contains returns true if the path contains a Data equal to the one specified
func (p *Path) Contains(key data.Data) bool {
	for _, d := range p.keys {
		if data.Equals(d, key) {
			return true
		}
	}
	return false
}
