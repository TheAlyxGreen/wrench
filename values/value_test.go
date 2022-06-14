package values

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestValue_Length(t *testing.T) {

	// inputs
	inputBytes := []byte{12, 54, 18, 243}

	// expectations
	const expectedOutput = 4

	// begin inputBytes
	testValue := NewFromBytes(inputBytes)
	fmt.Printf(indent(2)+"Expected:\t%s\n", expectedOutput)
	fmt.Printf(indent(2)+"Got:\t\t%s\n", testValue.Length())
	if testValue.Length() != expectedOutput {
		t.Fatal("length of values bytes doesn't match length of input bytes")
	}
}

func TestValue_StringLength(t *testing.T) {

	// inputs
	const inputString = "this is a test"

	// expectations
	const expectedOutput = 14

	// begin test
	v := New(inputString)
	if v.StringLength() != expectedOutput {
		t.Fatal("length of new values from string don't match length of input string")
	}
}

func TestValue_Equals(t *testing.T) {

	// inputs
	const inputUint uint64 = 123456789123456789
	inputAsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(inputAsBytes, inputUint)

	// expectations
	expectedOutput := NewFromBytes(inputAsBytes)

	// begin inputUint
	valueBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(valueBytes, inputUint)
	testValue := NewFromBytes(valueBytes)
	if !testValue.Equals(expectedOutput) {
		t.Fatal("keys from same Uint do not Equal")
	}
}

func TestValue_String(t *testing.T) {

	// inputs
	const inputString = "this is a test"

	// expectations
	const expectedOutput = "this is a test"

	// begin test
	testValue := New(inputString)
	fmt.Printf(indent(2)+"Expected:\t%s\n", expectedOutput)
	fmt.Printf(indent(2)+"Got:\t\t%s\n", testValue.String())
	if testValue.String() != expectedOutput {
		t.Fatal("values from same string do not Equal")
	}

}
