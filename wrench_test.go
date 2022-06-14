package wrench

import (
	"testing"
)

func TestNewWrench(t *testing.T) {
	const testPathString = "/not/a/real/path"

	w, err := New(testDatabasePathString)
	if err != nil {
		if err == ErrorFileCannotBeFound {
			t.Fatal("New() throws ErrorFileCannotBeFound on valid File")
		} else {
			t.Fatal("New() throws unexpected error")
		}
	}
	if w.filePath != testDatabasePathString {
		t.Fatal("New() does not set File paths")
	}

	w, err = New(testPathString)
	if err == nil {
		t.Fatal("New() does not throw an error if File does not exist")
	} else if err != ErrorFileCannotBeFound {
		t.Fatal("New() throws unexpected error if File does not exist")
	}
}
