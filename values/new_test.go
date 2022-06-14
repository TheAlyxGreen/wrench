package values

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNewFromBytes(t *testing.T) {

	// inputs
	inputBytes := []byte{12, 54, 18, 243} // random

	// expectations
	expectedOutput := []byte{12, 54, 18, 243}

	// begin test
	fmt.Println(indent(1) + "Testing that bytes of generated value match input bytes...")
	testValue := NewFromBytes(inputBytes)
	fmt.Printf(indent(2)+"Expected:\t%v\n", expectedOutput)
	fmt.Printf(indent(2)+"Got:\t\t%v\n", testValue.bytes)
	if !bytes.Equal(testValue.bytes, expectedOutput) {
		t.Fatal("bytes of generated value don't match input bytes")
	}
}

func TestNew(t *testing.T) {

	// inputs
	const inputString = "this is a test"

	// expectations
	expectedOutput := []byte{116, 104, 105, 115, 32, 105, 115, 32, 97, 32, 116, 101, 115, 116}

	// begin test
	fmt.Println(indent(1) + "Testing that bytes of generated value match bytes of input string...")
	v := New(inputString)
	fmt.Printf(indent(2)+"Expected:\t%v\n", expectedOutput)
	fmt.Printf(indent(2)+"Got:\t\t%v\n", v.bytes)
	if !bytes.Equal(v.bytes, expectedOutput) {
		t.Fatal("bytes of new values from string don't match bytes of input string")
	}
}
