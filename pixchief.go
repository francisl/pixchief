package main

import (
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

func usage() {
	fmt.Println("Usage:  pixchief [source path] [destination path]")
}

func pathManagement(src, dst string) {
	if !checkDir(src, dst) {
		usage()
		return
	}

	fmt.Println(src)
	fmt.Println(dst)
	fmt.Println(checkDir(src, dst))
}

func main() {
	if len(os.Args) < 3 {
		usage()
		return
	}
	pathManagement(os.Args[1], os.Args[2])
}
