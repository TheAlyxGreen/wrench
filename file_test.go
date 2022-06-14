package wrench

import (
	"fmt"
	"testing"
)

func TestWrench_FilePath(t *testing.T) {
	const testPathString = "/not/a/real/path"
	w, _ := New(testPathString)
	w.filePath = testPathString

	if w.FilePath() != testPathString {
		fmt.Println(w.FilePath())
		fmt.Println(testPathString)
		t.Fatal("FilePath() does not equal input filepath")
	}
}

func Test_FileExists(t *testing.T) {
	if !fileExists("File.go") {
		t.Fatal("fileExists() returns false on File.go")
	}
}

func TestWrench_FilePathExists(t *testing.T) {
	const testPathString = "/not/a/real/path"
	w, _ := New(testPathString)
	if w.FilePathExists() {
		t.Fatal("FilePathExists() returns true on fake File")
	}
	w.filePath = "File.go"
	if !w.FilePathExists() {
		t.Fatal("FilePathExists() returns false on File.go")
	}
}
