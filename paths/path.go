package paths

import (
	"github.com/TheAlyxGreen/wrench/values"
)

// A Path is the series of buckets that must be traversed to arrive at a specific key
type Path struct {
	keys []values.Value
}

// String returns the paths as a parseable paths string
func (p Path) String() string {
	if len(p.keys) == 0 {
		return rootPathString
	}
	var pathString string
	for i, k := range p.keys {
		if i == 0 {
			pathString = rootPathString + k.String()
		} else {
			pathString = pathString + pathDividerString + k.String()
		}
	}
	return pathString
}

// Equals returns true if both paths contain the same keys in the same order
func (p Path) Equals(p2 Path) bool {
	if len(p.keys) != len(p2.keys) {
		return false
	} else if len(p.keys) == 0 {
		return true
	}
	for i, key := range p.keys {
		if !key.Equals(p2.keys[i]) {
			return false
		}
	}
	return true
}

// Contains returns true if the paths contains a Value equal to the one specified
func (p Path) Contains(key values.Value) bool {
	for _, d := range p.keys {
		if d.Equals(key) {
			return true
		}
	}
	return false
}

// Length returns the number of keys in the path
func (p Path) Length() int {
	return len(p.keys)
}
