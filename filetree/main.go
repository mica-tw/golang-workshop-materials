package main

import (
	"filetree/tree"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

func main() {
	root := "."

	fileTree := tree.Tree{
		FileName: root,
	}
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		fileTree.AddToTree(strings.Split(path, "/"))

		// fmt.Printf("tree after adding '%s' to it: %v\n", path, fileTree)

		return nil
	})

	fmt.Println(fileTree)

}
