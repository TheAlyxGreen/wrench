package values

// NewFromBytes returns a Value from the specified bytes
func NewFromBytes(value []byte) Value {
	out := Value{}
	for i := 0; i < len(value); i++ {
		out.bytes = append(out.bytes, value[i])
	}
	return out
}

// New returns a Value from the specified string
func New(value string) Value {
	return NewFromBytes([]byte(value))
}
