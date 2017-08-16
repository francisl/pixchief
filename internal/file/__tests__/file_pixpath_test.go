package file_test

import (
	"testing"

	"github.com/francisl/pixchief/internal/file"
)

func TestPixpath_checkdirShouldReturnAnError_whenNoSourceFile(t *testing.T) {

	checkErr := file.ArePathsValid("./baddir", "./tmp")
	if checkErr == nil {
		t.Errorf("Return true for and existing directory")
	}
}
