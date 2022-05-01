package data

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestNewData(t *testing.T) {
	test := []byte{12, 54, 18, 243}
	d := NewData(test)
	if !bytes.Equal(d.bytes, test) {
		t.Fatal("bytes of new data from bytes don't match input bytes")
	}
}

func Test_DataFromString(t *testing.T) {
	testString := "this is a test"
	d := NewDataFromString(testString)
	if !bytes.Equal(d.bytes, []byte(testString)) {
		t.Fatal("bytes of new data from string don't match bytes of input string")
	}
}

func Test_StringLength(t *testing.T) {
	testString := "this is a test"
	d := NewDataFromString(testString)
	if len(testString) != StringLength(d) {
		t.Fatal("length of new data from string don't match length of input string")
	}
}

func Test_Length(t *testing.T) {
	test := []byte{12, 54, 18, 243}
	d := NewData(test)
	if Length(d) != len(test) {
		t.Fatal("length of data bytes doesn't match length of input bytes")
	}
}

func Test_ToBytes(t *testing.T) {
	test := []byte{12, 54, 18, 243}
	d := NewData(test)
	if !bytes.Equal(ToBytes(d), test) {
		t.Fatal("data ToBytes doesn't match input bytes")
	}
}

func Test_ToString(t *testing.T) {
	test := "this is a test"
	d := NewDataFromString(test)
	if ToString(d) != test {
		t.Fatal("data ToString doesn't match input string")
	}
}

func Test_ToKeyString(t *testing.T) {
	test := "this/is\\a/test"
	result := "this\\/is\\\\a\\/test"
	d := NewDataFromString(test)
	if ToKey(d) != result {
		fmt.Println(ToKey(d))
		fmt.Println(result)
		t.Fatal("data ToKey doesn't escape properly")
	}
}

func Test_ToUint(t *testing.T) {
	test := uint64(12345678917181920)
	testAsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(testAsBytes, test)
	d := NewData(testAsBytes)
	if ToUint(d) != test {
		t.Fatal("data ToUint doesn't match input Uint")
	}
}

func Test_Equals(t *testing.T) {
	test := uint64(12345678917181920)
	testAsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(testAsBytes, test)
	d1 := NewData(testAsBytes)
	d2 := NewData(testAsBytes)
	if Equals(d1, d2) {
		t.Fatal("keys from same Uint do not Equal")
	}
}
