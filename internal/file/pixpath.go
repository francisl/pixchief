package file

import (
	"errors"
	"fmt"
	"os"
	"path"
)

// checkDir validate that both directory exist
// When source is missing, it return false
// when dst dont exist, it will create the directories
func checkDir(src, dst string) bool {
	if _, err := os.Stat(src); err != nil {
		return false
	}

	if err := CheckOrCreatePath(dst); err != nil {
		return false
	}

	return true
}

//ArePathsValid Validate source and destination path and create destination path
func ArePathsValid(src, dst string) error {
	if !checkDir(src, dst) {
		return errors.New("Invalid Path")
	}

	checkDir(src, dst)
	return nil
}

// CheckOrCreatePath will create the path if not exist
func CheckOrCreatePath(dstPath string) error {
	if _, err := os.Stat(dstPath); err != nil {
		if path.IsAbs(dstPath) {
			println("Creating path : ", dstPath)
			if err := os.MkdirAll(dstPath, os.ModePerm); err == nil {
				return nil
			}
		}
		return fmt.Errorf(fmt.Sprintf("Error Creating Directory : %s ", dstPath))
	}
	return nil

}
