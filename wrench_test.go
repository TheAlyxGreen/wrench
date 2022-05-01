package wrench

import (
	"testing"
)

func Test_NewWrench(t *testing.T) {
	_, err := NewWrench("/notARealFile")
	if err.Error() != ErrorFileDoesNotExist {
		t.Fatal("NewWrench does not throw correct error if file does not exist")
	}
}
