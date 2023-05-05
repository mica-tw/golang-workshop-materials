package tree

import (
	"testing"
)

func TestAddToTree(t *testing.T) {
	// Create a new tree
	tr := Tree{
		FileName: "root",
	}
	tr.AddToTree([]string{"dir", "subdir", "file.txt"}, 100)

	// Verify the size and filename of the root node
	if tr.FileName != "root" {
		t.Errorf("Expected root node filename to be 'root', but got %s", tr.FileName)
	}
	if tr.size != 0 {
		t.Errorf("Expected root node size to be 0, but got %d", tr.size)
	}

	// Verify the size and filename of the subdir node
	if tr.children[0].FileName != "dir" {
		t.Errorf("Expected subdir node filename to be 'dir', but got %s", tr.children[0].FileName)
	}
	if tr.children[0].size != 100 {
		t.Errorf("Expected subdir node size to be 100, but got %d", tr.children[0].size)
	}

	// Verify the size and filename of the file.txt node
	if tr.children[0].children[0].FileName != "subdir" {
		t.Errorf("Expected file.txt node filename to be 'subdir', but got %s", tr.children[0].children[0].FileName)
	}
	if tr.children[0].children[0].size != 100 {
		t.Errorf("Expected file.txt node size to be 100, but got %d", tr.children[0].children[0].size)
	}

	// Add another node with the same name as subdir
	tr.AddToTree([]string{"dir", "subdir2", "file2.txt"}, 200)

	// Verify that the existing subdir node has been updated with a new child
	if len(tr.children[0].children) != 2 {
		t.Errorf("Expected subdir node to have 2 children, but got %d", len(tr.children[0].children))
	}
}

func TestString(t *testing.T) {
	t.Run("normal case", func(t *testing.T) {
		// Create a new tree
		tr := Tree{
			FileName: "root",
		}
		tr.AddToTree([]string{"dir", "subdir", "file.txt"}, 100)
		tr.AddToTree([]string{"dir", "subdir", "file2.txt"}, 200)
		tr.AddToTree([]string{"dir", "subdir2", "file.txt"}, 300)
		tr.AddToTree([]string{"dir", "file.txt"}, 400)

		// Generate the string representation of the tree
		str := tr.String()

		// Verify the string representation of the tree
		expectedStr := `root
└-- dir
    ├-- subdir
    |   ├-- file.txt (100)
    |   └-- file2.txt (200)
    ├-- subdir2
    |   └-- file.txt (300)
    └-- file.txt (400)
`
		if str != expectedStr {
			t.Errorf("Expected tree string representation:\n%s\n\nGot:\n%s", expectedStr, str)
		}
	})

	t.Run("empty tree", func(t *testing.T) {
		var tr Tree

		if s := tr.String(); s == "" {
			t.Errorf("representation of empty tree should be empty string. Got %s", s)
		}
	})
}
