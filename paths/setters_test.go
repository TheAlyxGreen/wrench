package paths

import (
	"fmt"
	"testing"

	"github.com/TheAlyxGreen/wrench/values"
)

func TestAppend(t *testing.T) {

	// inputs
	const pathString = "/not/a/real"
	const keyToPush = "path"

	// expectations
	const expectedOutput = "/not/a/real/path"

	// begin test
	pathBeforePush, _ := New(pathString)
	pathAfterPush := Append(pathBeforePush, values.New(keyToPush))

	fmt.Println(indent(1) + "Testing that Append appends key correctly...")
	fmt.Printf(indent(2)+"Expected:\t%s\n", expectedOutput)
	fmt.Printf(indent(2)+"Got:\t\t%s\n", pathAfterPush)
	if pathAfterPush.String() != expectedOutput {
		t.Fatal("Append does not correctly append new key to Path")
	}

}

func TestAppendPath(t *testing.T) {

	// inputs
	const pathString1 = "/not/a"
	const pathString2 = "/real/path"

	// expectations
	const expectedOutput = "/not/a/real/path"

	// begin test
	path1, _ := New(pathString1)
	path2, _ := New(pathString2)
	pathAfterAppend := AppendPath(path1, path2)

	fmt.Println(indent(1) + "Testing that AppendPath appends path correctly...")
	fmt.Printf(indent(2)+"Expected:\t%s\n", expectedOutput)
	fmt.Printf(indent(2)+"Got:\t\t%s\n", pathAfterAppend)
	if pathAfterAppend.String() != expectedOutput {
		t.Fatal("Path is not appended correctly")
	}
}

func TestAppendKeys(t *testing.T) {

	// inputs
	const pathString = "/not/a"
	var keysToAppend = []values.Value{
		values.New("real"),
		values.New("path"),
	}

	// expectations
	const expectedOutput = "/not/a/real/path"

	// begin test
	path1, _ := New(pathString)

	pathAfterAppend := AppendKeys(path1, keysToAppend)

	fmt.Println(indent(1) + "Testing that AppendKeys appends correctly...")
	fmt.Printf(indent(2)+"Expected:\t%s\n", expectedOutput)
	fmt.Printf(indent(2)+"Got:\t\t%s\n", pathAfterAppend)
	if pathAfterAppend.String() != expectedOutput {
		t.Fatal("Path is not appended correctly")
	}
}
