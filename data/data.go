package data

import (
	"bytes"
)

// A Data is used to store a key or value in the database
type Data struct {
	bytes []byte
}

// Equals returns true if the values of the bytes of each Data are equal
func Equals(d1, d2 Data) bool {
	return bytes.Equal(d1.bytes, d2.bytes)
}
