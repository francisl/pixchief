package pixpath

import (
	"errors"
	"fmt"
	"os"
	"path"
)

func checkDir(src, dst string) bool {
	srcExist := false
	dstExist := false

	if _, err := os.Stat(src); err == nil {
		srcExist = true
	}

	if _, err := os.Stat(dst); err == nil {
		dstExist = true
	} else {
		if path.IsAbs(dst) {
			if err := os.MkdirAll(dst, os.ModePerm); err == nil {
				dstExist = true
			}
		}
	}

	return srcExist && dstExist
}

//PathManagement Validate source and destination path and create destination path
func PathManagement(src, dst string) error {
	if !checkDir(src, dst) {
		return errors.New("Invalid Path")
	}

	fmt.Println(src)
	fmt.Println(dst)
	fmt.Println(checkDir(src, dst))

	return nil
}
