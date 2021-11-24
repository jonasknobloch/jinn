package tree

func (t *Tree) Copy(m map[*Tree]*Tree) *Tree {
	if t == nil {
		return nil
	}

	var children []*Tree

	if len(t.Children) > 0 {
		children = make([]*Tree, len(t.Children))

		for k, c := range t.Children {
			children[k] = c.Copy(m)
		}
	}

	cp := &Tree{
		Label:    t.Label,
		Children: children,
	}

	if m != nil {
		m[cp] = t
	}

	return cp
}
