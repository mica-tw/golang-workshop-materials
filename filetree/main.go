package main

import (
	"filetree/tree"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var root string

	switch i := len(os.Args); i {
	case 1:
		root = "."
	case 2:
		root = os.Args[1]
	default:
		fmt.Printf("You supplied %d arguments, a maximum of 1 is allowed", i-1)
		os.Exit(1)
	}

	fileTree := tree.Tree{
		FileName: root,
	}
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		fileTree.AddToTree(strings.Split(path, "/"))

		return nil
	})

	fmt.Println(fileTree)

}
