package main

import (
	"fmt"
	"os"

	"github.com/francisl/pixchief/internal/pixpath"
)

func usage() {
	fmt.Println("Usage:  pixchief [source path] [destination path]")
}

func main() {
	if len(os.Args) < 3 {
		usage()
		return
	}
	err := pixpath.PathManagement(os.Args[1], os.Args[2])
	if err != nil {
		usage()
	}
}
