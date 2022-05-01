package path

import (
	"github.com/TheAlyxGreen/wrench/data"
)

// PushKey appends the given Data to the end of the path
func PushKey(p Path, key data.Data) Path {
	p.keys = append(p.keys, key)
	return p
}

// AppendPath appends all the keys from one path onto another path
func AppendPath(p Path, p2 Path) Path {
	for _, key := range GetKeys(p2) {
		p.keys = append(p.keys, key)
	}
	return p
}

// AppendKeys appends an array of keys onto the path
func AppendKeys(p Path, keys []data.Data) Path {
	for _, key := range keys {
		p.keys = append(p.keys, key)
	}
	return p
}
