package values

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestValue_GetBytes(t *testing.T) {

	// inputs
	inputBytes := []byte{12, 54, 18, 243}

	// begin test
	fmt.Println(indent(1) + "Testing that output bytes match input bytes...")
	v := NewFromBytes(inputBytes)
	fmt.Printf(indent(2)+"Expected:\t%v\n", inputBytes)
	fmt.Printf(indent(2)+"Got:\t\t%v\n", v.GetBytes())
	if !bytes.Equal(v.GetBytes(), inputBytes) {
		t.Fatal("values GetBytes doesn't match input bytes")
	}
}

func TestValue_GetRawString(t *testing.T) {

	// inputs
	const inputString = "this is a test"

	// begin test
	fmt.Println(indent(1) + "Testing that output string matches input string...")
	v := New(inputString)
	fmt.Printf(indent(2)+"Expected:\t%v\n", inputString)
	fmt.Printf(indent(2)+"Got:\t\t%v\n", v.GetRawString())
	if v.GetRawString() != inputString {
		t.Fatal("output string doesn't match input string")
	}
}

func TestValue_GetUint(t *testing.T) {

	// inputs
	const inputUint uint64 = 123456789123456789
	inputAsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(inputAsBytes, inputUint)

	// begin test
	fmt.Println(indent(1) + "Testing that the output uint64 is the same as the input uint64...")
	testValue := NewFromBytes(inputAsBytes)
	fmt.Printf(indent(2)+"Expected:\t%d\n", inputUint)
	fmt.Printf(indent(2)+"Got:\t\t%d\n", testValue.GetUint())
	if testValue.GetUint() != inputUint {
		t.Fatal("values GetUint doesn't match input Uint")
	}
}

func TestValue_GetPathKey(t *testing.T) {

	// inputs
	const inputString = "this/is\\a/test"

	// expectations
	const expectedOutput = "this\\/is\\\\a\\/test"

	// begin test
	fmt.Println(indent(1) + "Testing that input string escaped properly...")
	testValue := New(inputString)
	fmt.Printf(indent(2)+"Expected:\t%s\n", expectedOutput)
	fmt.Printf(indent(2)+"Got:\t\t%s\n", testValue.GetPathKey())
	if testValue.GetPathKey() != expectedOutput {
		t.Fatal("output string is incorrect")
	}
}
