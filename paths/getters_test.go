package paths

import (
	"fmt"
	"github.com/TheAlyxGreen/wrench/values"
	"testing"
)

func TestPath_Keys(t *testing.T) {
	// inputs
	const inputString = "/not/a/real/path"

	// expectations
	expectedOutput := []values.Value{
		values.New("not"),
		values.New("a"),
		values.New("real"),
		values.New("path"),
	}

	// begin test
	testPath, _ := New(inputString)

	fmt.Println(indent(1) + "Testing that Keys() returns correct value...")
	fmt.Printf(indent(2)+"Expected:\t%v\n", expectedOutput)
	fmt.Printf(indent(2)+"Got:\t\t%v\n", testPath.Keys())
	if len(testPath.keys) != len(testPath.Keys()) {
		t.Fatal("returns wrong number of keys")
	}
	keys := testPath.Keys()
	for i, key := range testPath.keys {
		if !keys[i].Equals(key) {
			fmt.Printf(indent(2)+"Expected:\t%s\n", key.GetRawString())
			fmt.Printf(indent(2)+"Got:\t\t%s\n", keys[i].GetRawString())
			t.Fatal("key mismatch")
		}
	}
}

func TestPeekKey(t *testing.T) {
	// inputs
	const inputString = "/not/a/real/path"

	// expectations
	const expectedLastKey = "path"

	// begin test
	testPath1, _ := New(inputString)

	fmt.Println(indent(1) + "Testing that PeekKey shows last key...")
	returnedKey, err1 := testPath1.PeekKey()

	if err1 != err1 {
		fmt.Println(err1)
		t.Fatal("throws error on test path")
	}

	fmt.Printf(indent(2)+"Expected:\t%s\n", expectedLastKey)
	fmt.Printf(indent(2)+"Got:\t\t%s\n", returnedKey.GetRawString())
	if returnedKey.GetRawString() != expectedLastKey {
		t.Fatal("returns incorrect key")
	}

	fmt.Println(indent(1) + "Testing that PeekKey throws error on blank path...")

	testPath2 := Path{}
	_, err2 := testPath2.PeekKey()

	if err2 == nil {
		t.Fatal("does not throw error on blank path")
	} else if err2 != ErrorPathHasNoKeys {
		t.Fatal("throws incorrect error on blank path")
	}

}

func TestPopKey(t *testing.T) {
	// inputs
	const inputString = "/not/a/real/path"

	// expectations
	const expectedLastKey = "path"
	const expectedPathAfterPop = "/not/a/real"

	// begin test
	testPath1, _ := New(inputString)

	fmt.Println(indent(1) + "Testing that PopKey shows last key...")
	returnedKey, err1 := testPath1.PopKey()

	if err1 != err1 {
		fmt.Println(err1)
		t.Fatal("throws error on test path")
	}

	fmt.Printf(indent(2)+"Expected:\t%s\n", expectedLastKey)
	fmt.Printf(indent(2)+"Got:\t\t%s\n", returnedKey.GetRawString())
	if returnedKey.GetRawString() != expectedLastKey {
		t.Fatal("returns wrong key")
	}

	fmt.Println(indent(1) + "Testing that PopKey correctly removes last key...")
	fmt.Printf(indent(2)+"Expected:\t%s\n", expectedPathAfterPop)
	fmt.Printf(indent(2)+"Got:\t\t%s\n", testPath1)
	if testPath1.String() != expectedPathAfterPop {
		t.Fatal("does not correctly remove final key")
	}

	fmt.Println(indent(1) + "Testing that PopKey throws error on blank path...")

	testPath2 := Path{}
	_, err2 := testPath2.PopKey()

	if err2 == nil {
		t.Fatal("does not throw error on blank path")
	} else if err2 != ErrorPathHasNoKeys {
		t.Fatal("throws incorrect error on blank path")
	}

}
