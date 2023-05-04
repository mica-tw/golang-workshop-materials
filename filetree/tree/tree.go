package tree

type Tree map[string]Tree

func (t Tree) AddToTree(s []string) {
	if len(s) == 0 {
		return
	}

	if t == nil {
		t = make(Tree)
	}

	firstElem := s[0]

	if len(s) > 1 {
		if nextBranch, ok := t[firstElem]; ok {
			nextBranch.AddToTree(s[1:])
			return
		}

		t[firstElem] = createTreeFromPath(s[1:])
		return
	}

	if _, ok := t[firstElem]; !ok {
		t[firstElem] = nil
	}
}

func createTreeFromPath(s []string) Tree {
	if len(s) == 1 {
		return Tree{s[0]: nil}
	}

	return Tree{s[0]: createTreeFromPath(s[1:])}
}
