package tree

func (t *Tree) Equals(e *Tree) bool {
	if t.Label != e.Label {
		return false
	}

	if len(t.Children) != len(e.Children) {
		return false
	}

	for k, v := range t.Children {
		if v.Equals(e.Children[k]) {
			continue
		}

		return false
	}

	return true
}
