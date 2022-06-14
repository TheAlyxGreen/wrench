package values

import (
	"encoding/binary"
)

// GetBytes returns the raw bytes of the Value
func (v Value) GetBytes() []byte {
	return v.bytes
}

// GetRawString returns the bytes of the Value as a string
func (v Value) GetRawString() string {
	return string(v.bytes)
}

// GetPathKey returns the Value with any slashes escaped to use in a path
func (v Value) GetPathKey() string {
	stringValue := string(v.bytes)
	var newString string
	for _, c := range stringValue {
		currentChar := string(c)
		switch currentChar {
		case "\\":
			newString = newString + "\\" + currentChar
			break
		case "/":
			newString = newString + "\\" + currentChar
			break
		default:
			newString = newString + currentChar
		}
	}
	return newString
}

// GetUint returns the bytes of the Value as an unsigned int64
func (v Value) GetUint() uint64 {
	return binary.LittleEndian.Uint64(v.bytes)
}
