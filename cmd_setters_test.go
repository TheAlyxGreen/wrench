package wrench

import (
	"fmt"
	"testing"

	"github.com/TheAlyxGreen/wrench/paths"
	"github.com/TheAlyxGreen/wrench/values"
)

func TestWrench_InsertBucket(t *testing.T) {
	// inputs
	testPath, _ := paths.New("/CreateBucket")
	w, _ := New(testDatabasePathString)

	// begin test
	fmt.Println(indent(1) + "Testing that it fails when passed root bucket path...")
	err := w.CreateBucket(paths.RootPath)
	if err == nil {
		fmt.Printf(indent(2)+"Expected: %s\n", err)
		fmt.Printf(indent(2)+"Got: %s\n", w.ErrMsg())
		t.Fatalf("does not throw error on root path")
	} else if err != ErrorPathCannotBeRoot {
		fmt.Printf(indent(2)+"Expected: %s\n", err)
		fmt.Printf(indent(2)+"Got: %s\n", w.ErrMsg())
		t.Fatalf("throws incorrect error on root path")
	}

	fmt.Println(indent(1) + "Testing that it works on valid path...")
	err = w.CreateBucket(testPath)
	if err != nil {
		fmt.Printf(indent(2)+"Expected: %s\n", err)
		fmt.Printf(indent(2)+"Got: \n%s\n", w.ErrMsg())
		t.Fatalf("throws error on valid path: %v", err)
	}

	err = resetTestDatabase()
	if err != nil {
		t.Fatal("could not reset database after test")
	}
}

func TestWrench_SetValue(t *testing.T) {

	// inputs
	testPath, _ := paths.New(testValuePathString)
	fakePath, _ := paths.New("/not/a/real/path")

	// begin test

	w, err := New("./testing/test.db")
	if err != nil {
		t.Fatal(err)
	}

	err = w.SetValue(fakePath, values.New("test"))
	if err == nil {
		t.Fatal("does not throw error on non-existent path")
	} else if err != ErrorPathCannotBeFound {
		t.Fatal("throws incorrect error on non-existent path")
	}

	err = w.SetValue(testPath, values.New("test"))
	if err != nil {
		t.Fatal("throws error on valid path")
	}

	fmt.Println("Incomplete test: does not test db to see if value was set correctly")
	t.Fail()

	err = resetTestDatabase()
	if err != nil {
		t.Fatal("could not reset database after test")
	}

}
