package data

import (
	"encoding/binary"
	"time"
)

// ToBytes returns the raw bytes of the Data
func ToBytes(d Data) []byte {
	return d.bytes
}

// ToString returns the bytes of the Data as a string
func ToString(d Data) string {
	return string(d.bytes)
}

// ToKey returns the Data with any slashes escaped to use as a key
func ToKey(d Data) string {
	stringValue := string(d.bytes)
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

// ToUint returns the bytes of the Data as an unsigned int64
func ToUint(d Data) uint64 {
	return binary.LittleEndian.Uint64(d.bytes)
}

func ToTime(d Data, layout string) (time.Time, error) {
	result, err := time.Parse(layout, ToString(d))
	if err != nil {
		return time.Time{}, err
	}
	return result, nil
}
