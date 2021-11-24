package tree

import (
	"strings"
)

type Tree struct {
	Label    string
	Children []*Tree
}

func (t *Tree) Walk(cb func(t *Tree)) {
	cb(t)

	for _, c := range t.Children {
		c.Walk(cb)
	}
}

func (t *Tree) WalkPositions(cb func(t *Tree, p []int), p []int) {
	cb(t, p)

	for s, c := range t.Children {
		c.WalkPositions(cb, append(p, s))
	}
}

func (t *Tree) Height() int {
	max := 0

	for _, p := range t.Positions() {
		if max < len(p)+1 {
			max = len(p) + 1
		}
	}

	return max
}

func (t *Tree) Size() int {
	return len(t.Positions())
}

func (t *Tree) Positions() [][]int {
	ps := make([][]int, 0)

	t.WalkPositions(func(t *Tree, p []int) {
		ps = append(ps, p)
	}, []int{})

	return ps
}

func (t *Tree) Subtrees() []*Tree {
	sts := make([]*Tree, 0)

	t.Walk(func(t *Tree) {
		sts = append(sts, t)
	})

	return sts
}

func (t *Tree) SubtreeAtPosition(p []int) *Tree {
	if len(p) == 0 {
		return t
	}

	return t.Children[p[0]].SubtreeAtPosition(p[1:])
}

func (t *Tree) ReplaceAtPosition(p []int, r *Tree) {
	if len(p) == 1 {
		t.Children[p[0]] = r
	}

	t.ReplaceAtPosition(p[1:], r)
}

func (t *Tree) LabelAtPosition(p []int) string {
	if len(p) == 0 {
		return t.Label
	}

	return t.Children[p[0]].LabelAtPosition(p[1:])
}

func (t *Tree) Leaves() []*Tree {
	ls := make([]*Tree, 0)

	t.Walk(func(t *Tree) {
		if len(t.Children) == 0 {
			ls = append(ls, t)
		}
	})

	return ls
}

func (t *Tree) LeafPositions() [][]int {
	lps := make([][]int, 0)

	t.WalkPositions(func(t *Tree, p []int) {
		if len(t.Children) == 0 {
			lps = append(lps, p)
		}
	}, []int{})

	return lps
}

func (t *Tree) Edges() map[*Tree][]*Tree {
	es := make(map[*Tree][]*Tree)

	t.Walk(func(t *Tree) {
		if len(t.Children) != 0 {
			es[t] = t.Children
		}
	})

	return es
}

func (t *Tree) EdgePositions() map[string][]string {
	eps := make(map[string][]string)

	t.WalkPositions(func(t *Tree, p []int) {
		if len(t.Children) == 0 {
			return
		}

		children := make([]string, 0)

		for s := range t.Children {
			children = append(children, Position(append(p, s)).String())
		}

		eps[Position(p).String()] = children
	}, []int{})

	return eps
}

func (t *Tree) Sentence() string {
	var sb strings.Builder

	for _, l := range t.Leaves() {
		sb.WriteString(" ")
		sb.WriteString(l.Label)
	}

	return sb.String()[1:]
}

func (t *Tree) String() string {
	var sb strings.Builder

	var walk func(t *Tree)
	walk = func(t *Tree) {
		sb.WriteString(" ")

		if len(t.Children) == 0 {
			sb.WriteString(t.Label)
			return
		}

		sb.WriteString("(")
		sb.WriteString(t.Label)

		for _, c := range t.Children {
			walk(c)
		}

		sb.WriteString(")")
	}

	walk(t)

	return sb.String()[1:]
}
