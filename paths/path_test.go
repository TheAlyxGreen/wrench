package paths

import (
	"fmt"
	"testing"

	"github.com/TheAlyxGreen/wrench/values"
)

func TestPath_String(t *testing.T) {

	// inputs
	const inputString = "/not/a/real/path"

	// begin test
	testPath, _ := New(inputString)

	fmt.Println(indent(1) + "Testing that Path.String() matches input string...")
	fmt.Printf(indent(2)+"Expected:\t%s\n", inputString)
	fmt.Printf(indent(2)+"Got:\t\t%s\n", testPath.String())
	if testPath.String() != inputString {
		t.Fatal("generated Path string does not match input")
	}
	testPath = Path{}
	fmt.Println(indent(1) + "Testing that Path.String() returns root when there are no keys...")
	fmt.Printf(indent(2)+"Expected:\t%s\n", rootPathString)
	fmt.Printf(indent(2)+"Got:\t\t%s\n", testPath.String())
	if testPath.String() != rootPathString {
		t.Fatal("path does not return rootPathString when there are no keys")
	}

}

func TestPath_Equals(t *testing.T) {

	// inputs
	const inputString1 = "/not/a/real/path"
	testPath, _ := New(inputString1)
	testPathCopy, _ := New(inputString1)

	const inputString2 = "/this/is/a/test" // same length as inputString1 but different keys
	testPath2, _ := New(inputString2)

	shortPath, _ := New("/first/second")      // all keys from short path
	longPath, _ := New("/first/second/third") // are in long path

	// begin test
	fmt.Println(indent(1) + "Testing if a Path equals itself...")
	if !testPath.Equals(testPath) {
		t.Fatal("generated Path does not equal itself")
	}

	fmt.Println(indent(1) + "Testing if Path equals another Path from same string...")

	if !testPath.Equals(testPathCopy) {
		t.Fatal("generated Path does not equal second Path generated from same string")
	}

	fmt.Println(indent(1) + "Testing mismatched paths...")
	fmt.Printf(indent(2)+"Path 1:\t%s\n", testPath)
	fmt.Printf(indent(2)+"Path 2:\t%s\n", testPath2)
	// check that paths don't match if their keys don't match
	if testPath.Equals(testPath2) {
		t.Fatal("returned true for mismatched paths of equal length")
	}

	// check that paths with different lengths don't match, even if the first parts do
	fmt.Println(indent(1) + "Testing partially-matched paths...")
	fmt.Printf(indent(2)+"Path 1:\t%s\n", shortPath)
	fmt.Printf(indent(2)+"Path 2:\t%s\n", longPath)
	if shortPath.Equals(longPath) {
		t.Fatal("returned true for mismatched Paths of unequal lengths")
	}

}

func TestPath_Contains(t *testing.T) {

	// inputs
	const inputString = "/not/a/real/path"
	const realKey = "not"
	const nonexistentKey = "doesNotExist"

	// begin test
	testPath, _ := New(inputString)
	fmt.Println(indent(1) + "Testing with key that is contained in the Path...")
	fmt.Printf(indent(2)+"Path:\t%s\n", testPath)
	fmt.Printf(indent(2)+"Key:\t%s\n", realKey)
	if !testPath.Contains(values.New(realKey)) {
		t.Fatal("returns false on key that is contained in the Path")
	}

	fmt.Println(indent(1) + "Testing with a key that is not contained in the Path...")
	fmt.Printf(indent(2)+"Path:\t%s\n", testPath)
	fmt.Printf(indent(2)+"Key:\t%s\n", nonexistentKey)
	if testPath.Contains(values.New(nonexistentKey)) {
		t.Fatal("returns true on key that is not contained in the Path")
	}
}

func TestPath_Length(t *testing.T) {

	// inputs
	const inputString = "/not/a/real/path"

	// expectations
	const expectedKeyCount = 4

	// begin test
	testPath, _ := New(inputString)
	fmt.Println(indent(1) + "Testing that Path reports correct length...")
	fmt.Printf(indent(2)+"Expected:\t%d\n", expectedKeyCount)
	fmt.Printf(indent(2)+"Got:\t\t%d\n", testPath.Length())
	if testPath.Length() != expectedKeyCount {
		t.Fatal("returns the wrong number")
	}
}
