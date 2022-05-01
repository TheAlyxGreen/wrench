package data

// NewData returns a Data from the specified bytes
func NewData(value []byte) Data {
	return Data{bytes: value}
}

// NewDataFromString returns a Data from the specified string
func NewDataFromString(value string) Data {
	return NewData([]byte(value))
}
