package values

import (
	"bytes"
)

// A Value is used to store a key or value in the database. Assumed to be a string
type Value struct {
	bytes []byte
}

// Length returns the length of the Value as an array of bytes
func (v Value) Length() int {
	return len(v.bytes)
}

// StringLength parses the Value as a string then returns the length of that string
func (v Value) StringLength() int {
	return len(string(v.bytes))
}

// Equals returns true if the values of the bytes of each Value are equal
func (v Value) Equals(d2 Value) bool {
	return bytes.Equal(v.bytes, d2.bytes)
}

// String returns the Value as a string
func (v Value) String() string {
	return string(v.bytes)
}
