package main

import (
	"fmt"
	"os"

	"github.com/francisl/pixchief/internal/file"
)

func usage() {
	fmt.Println("Usage:  pixchief [source path] [destination path]")
}

func main() {
	if len(os.Args) < 3 {
		usage()
		return
	}
	src, dest := os.Args[1], os.Args[2]
	err := file.ArePathsValid(src, dest)
	if err != nil {
		usage()
	}
	file.MoveFiles(src, dest)
}
