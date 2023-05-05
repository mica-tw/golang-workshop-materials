package tree

import (
	"bytes"
	"fmt"
	"io"

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

func (t *Tree) recursiveString(w io.Writer, prefix string) error {
	s := t.FileName

	if len(t.children) == 0 {
		s = fmt.Sprintf("%s (%d)", s, t.size)
	}

	_, err := w.Write([]byte(fmt.Sprintf("%s\n", s)))
	if err != nil {
		return fmt.Errorf("failed to write tree: %w", err)
	}

	for i, child := range t.children {
		startingChar := "├-- "
		nextPrefix := fmt.Sprintf("%s%s", prefix, "|   ")
		if i == len(t.children)-1 {
			startingChar = "└-- "
			nextPrefix = fmt.Sprintf("%s%s", prefix, "    ")
		}

		w.Write(append([]byte(prefix), []byte(startingChar)...))
		if err != nil {
			return fmt.Errorf("failed to write tree: %w", err)
		}

		err = child.recursiveString(w, nextPrefix)
		if err != nil {
			return fmt.Errorf("could not recurse: %w", err)
		}
	}

	return nil
}

func (t Tree) String() string {
	var bufBytes []byte
	buf := bytes.NewBuffer(bufBytes)

	t.recursiveString(buf, "")

	res, err := io.ReadAll(buf)
	if err != nil {
		panic(err)
	}

	return string(res)
}
