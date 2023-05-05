package tree

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type Tree struct {
	FileName string
	size     int64
	children []Tree
}

func (t *Tree) AddToTree(s []string, size int64) {
	if len(s) == 0 {
		return
	}

	var nextIter []string
	if len(s) > 1 {
		nextIter = s[1:]
	}

	firstElem := s[0]

	matchesElem := func(tr Tree) bool {
		return tr.FileName == firstElem
	}

	if i := slices.IndexFunc(t.children, matchesElem); i != -1 {
		t.children[i].AddToTree(nextIter, size)
		return
	}

	newChild := Tree{
		FileName: firstElem,
		size:     size,
	}

	newChild.AddToTree(nextIter, size)

	t.children = append(t.children, newChild)
}

func (t *Tree) recursiveString(prefix string) string {
	s := t.FileName

	if len(t.children) == 0 {
		s = fmt.Sprintf("%s (%d)", s, t.size)
	}

	s = fmt.Sprintf("%s\n", s)

	for i, child := range t.children {
		startingChar := "├"
		nextPrefix := fmt.Sprintf("%s%s", prefix, "|   ")
		if i == len(t.children)-1 {
			startingChar = "└"
			nextPrefix = fmt.Sprintf("%s%s", prefix, "    ")
		}

		s = fmt.Sprintf("%s%s%s-- %s", s, prefix, startingChar,
			child.recursiveString(nextPrefix))
	}

	return s
}

func (t Tree) String() string {
	return t.recursiveString("")
}
