package data

// Length returns the length of the Data as an array of bytes
func Length(d Data) int {
	return len(d.bytes)
}

func StringLength(d Data) int {
	return len(string(d.bytes))
}
