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

	// filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
	// 	if path == root {
	// 		fmt.Println(root)
	// 	}

	// 		splitPath := strings.Split(path, "/")

	// 	toPrint := fmt.Sprintf("%s├── %s (%d)\n", strings.Repeat("|  ", len(splitPath)-1), splitPath[len(splitPath)-1], info.Size())

	// 	fmt.Print(toPrint)

	// 	return nil
	// })

	fileTree := make(tree.Tree)
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		fileTree.AddToTree(strings.Split(path, "/"))

		fmt.Printf("tree after adding '%s' to it: %v\n", path, fileTree)

		return nil
	})

	fmt.Println(fileTree)
}
