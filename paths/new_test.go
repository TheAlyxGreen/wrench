package paths

import (
	"fmt"
	"testing"

	"github.com/TheAlyxGreen/wrench/values"
)

func TestNewFromKeys(t *testing.T) {

	// inputs
	const inputString = "/not/a/real/path"

	// expectations
	const expectedKeyCount = 4
	expectedPathStrings := []string{"not", "a", "real", "path"}
	expectedPathValues := []values.Value{
		values.New("not"),
		values.New("a"),
		values.New("real"),
		values.New("path"),
	}

	// validate inputs
	if len(expectedPathStrings) != expectedKeyCount || len(expectedPathValues) != expectedKeyCount {
		t.Fatal("Test inputs are broken")
	}

	// begin test
	fmt.Println(indent(1) + "Creating Path from string with known path...")

	fmt.Printf(indent(2)+"Initial Path:\t%s\n", inputString)

	generatedPath := NewFromKeys(expectedPathValues)
	fmt.Printf(indent(2)+"Created Path:\t%s\n", generatedPath)

	fmt.Println(indent(1) + "Comparing the Paths...")

	fmt.Printf(indent(2)+"Expected:\t%s\n", expectedKeyCount)
	fmt.Printf(indent(2)+"Got:\t\t%s\n", len(generatedPath.keys))
	if len(expectedPathStrings) != len(generatedPath.keys) {
		t.Fatal("generated Path has wrong number of keys")
	}

	for i, v := range generatedPath.Keys() {
		if v.GetRawString() != expectedPathStrings[i] {
			t.Fatal("generated Path has wrong keys")
		}
	}
}

func TestNew(t *testing.T) {

	// inputs
	const inputString = "/not/a/real/path"

	// expectations
	const expectedKeyCount = 4
	expectedPathStrings := [4]string{"not", "a", "real", "path"}

	// begin test
	fmt.Println(indent(1) + "Creating Path from string...")
	fmt.Printf(indent(2)+"Initial string:\t%s\n", inputString)

	generatedPath, err := New(inputString)

	if err != nil {
		if err == ErrorInvalidPath {
			t.Fatal("threw parsing error")
		} else {
			t.Fatalf("threw unexpected error: %s\n", err.Error())
		}
	} else {
		fmt.Printf(indent(2)+"Created Path:\t\t%s\n", generatedPath.String())
	}

	fmt.Println(indent(1) + "Comparing the Path keys...")
	if len(generatedPath.keys) != expectedKeyCount {
		fmt.Printf(indent(2)+"Expected:\t%s\n", expectedKeyCount)
		fmt.Printf(indent(2)+"Got:\t\t%s\n", len(generatedPath.keys))
		t.Fatal("generated Path has wrong number of keys")
	}

	for i, v := range generatedPath.Keys() {
		if v.GetRawString() != expectedPathStrings[i] {
			fmt.Printf(indent(2)+"Expected:\t%s\n", expectedPathStrings[i])
			fmt.Printf(indent(2)+"Got:\t\t%s\n", v.GetRawString())
			t.Fatal("generated Path has wrong keys")
		}
	}

}

func TestNewPathFromEscapedString(t *testing.T) {

	// inputs
	const inputString = "/not/a\\/real/path"

	// expectations
	const expectedKeyCount = 3
	expectedPathStrings := [3]string{"not", "a/real", "path"}

	// begin test
	fmt.Println(indent(1) + "Creating Path from string...")
	fmt.Printf(indent(2)+"Initial string:\t%s\n", inputString)
	generatedPath, err := New(inputString)

	if err != nil {
		if err == ErrorInvalidPath {
			t.Fatal("threw parsing error")
		} else {
			t.Fatalf("threw unexpected error: %s\n", err.Error())
		}
	} else {
		fmt.Printf(indent(2)+"Created Path:\t\t%s\n", generatedPath.String())
	}

	fmt.Println(indent(1) + "Comparing the Path keys...")
	if len(generatedPath.keys) != expectedKeyCount {
		fmt.Printf(indent(2)+"Expected:\t%d\n", expectedKeyCount)
		fmt.Printf(indent(2)+"Got:\t\t%d\n", len(generatedPath.keys))
		t.Fatal("generated Path has wrong number of keys")
	}
	for i, v := range generatedPath.Keys() {
		if v.GetRawString() != expectedPathStrings[i] {
			fmt.Printf(indent(2)+"Expected:\t%s\n", expectedPathStrings[i])
			fmt.Printf(indent(2)+"Got:\t\t%s\n", v.GetRawString())
			t.Fatal("generated Path is wrong")
		}
	}
}

func TestNewPathFromWeirdStrings(t *testing.T) {

	// inputs
	const blankString = ""
	const invalidString = "//"

	// begin test
	fmt.Println(indent(1) + "Testing blank string...")

	_, err1 := New(blankString)
	fmt.Printf(indent(2)+"String: %s\n", blankString)
	if err1 == nil {
		t.Fatal("does not throw error on blank paths")
	} else if err1 != ErrorInvalidPath {
		t.Fatal("throws wrong error on blank paths")
	}

	fmt.Println(indent(1) + "Testing invalid string...")
	fmt.Printf(indent(2)+"String: %s\n", invalidString)
	_, err2 := New(invalidString)
	if err2 == nil {
		t.Fatal("does not throw error on invalid paths")
	} else if err2 != ErrorInvalidPath {
		t.Fatal("throws wrong error on invalid paths")
	}

	fmt.Println(indent(1) + "Testing rootPathString...")
	fmt.Printf(indent(2)+"String: %s\n", rootPathString)
	rootPathFromString, err3 := New(rootPathString)
	if err3 != nil {
		t.Fatal("throws error on root paths string")
	}
	if !rootPathFromString.Equals(RootPath) {
		t.Fatal("returns invalid paths on root paths string")
	}
}
