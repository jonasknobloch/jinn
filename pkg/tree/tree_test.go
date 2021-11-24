package tree

import "testing"

func TestTree_Height(t *testing.T) {
	p := NewDecoder()

	cases := []struct {
		tree   *Tree
		height int
	}{
		{
			p.MustDecode("(S (NP (NNP John)) (VP (V runs)) (. .))"),
			4,
		},
	}

	for _, c := range cases {
		s := c.tree.Height()

		if s != c.height {
			t.Errorf("%v expected but got %v\n", c.height, s)
		}
	}
}

func TestTree_Size(t *testing.T) {
	p := NewDecoder()

	cases := []struct {
		tree *Tree
		size int
	}{
		{
			p.MustDecode("(S (NP (NNP John)) (VP (V runs)) (. .))"),
			9,
		},
	}

	for _, c := range cases {
		s := c.tree.Size()

		if s != c.size {
			t.Errorf("%v expected but got %v\n", c.size, s)
		}
	}
}

func TestTree_Positions(t *testing.T) {
	p := NewDecoder()

	cases := []struct {
		tree      *Tree
		positions [][]int
	}{
		{
			p.MustDecode("(S (NP (NNP John)) (VP (V runs)) (. .))"),
			[][]int{{}, {0}, {0, 0}, {0, 0, 0}, {1}, {1, 0}, {1, 0, 0}, {2}, {2, 0}},
		},
	}

	equals := func(p1, p2 []int) bool {
		if len(p1) != len(p2) {
			return false
		}

		for k, v := range p1 {
			if p2[k] != v {
				return false
			}
		}

		return true
	}

	for _, c := range cases {
		ps := c.tree.Positions()

		if len(c.positions) != len(ps) {
			t.Errorf("%v expected but got %v\n", c.positions, ps)
		}

		for k, p := range c.positions {
			if !equals(ps[k], p) {
				t.Errorf("%v expected but got %v\n", p, ps[k])
			}
		}
	}
}

func TestTree_Subtrees(t *testing.T) {
	p000 := &Tree{"John", nil}
	p00 := &Tree{"NNP", []*Tree{p000}}
	p0 := &Tree{"NP", []*Tree{p00}}

	p100 := &Tree{"runs", nil}
	p10 := &Tree{"VP", []*Tree{p100}}
	p1 := &Tree{"V", []*Tree{p10}}

	p20 := &Tree{".", nil}
	p2 := &Tree{".", []*Tree{p20}}

	root := &Tree{
		"S", []*Tree{p0, p1, p2},
	}

	cases := []struct {
		tree     *Tree
		subtrees []*Tree
	}{
		{
			root,
			[]*Tree{root, p0, p00, p000, p1, p10, p100, p2, p20},
		},
	}

	for _, c := range cases {
		sts := c.tree.Subtrees()

		if len(c.subtrees) != len(sts) {
			t.Errorf("%v expected but got %v\n", c.subtrees, sts)
		}

		for k, st := range c.subtrees {
			if sts[k] != st {
				t.Errorf("%v expected but got %v\n", st, sts[k])
			}
		}
	}
}

func TestTree_SubtreeAtPosition(t *testing.T) {
	// TODO implement
}

func TestTree_ReplaceAtPosition(t *testing.T) {
	// TODO implement
}

func TestTree_LabelAtPosition(t *testing.T) {
	// TODO implement
}

func TestTree_Leaves(t *testing.T) {
	p000 := &Tree{"John", nil}
	p100 := &Tree{"runs", nil}
	p20 := &Tree{".", nil}

	cases := []struct {
		tree   *Tree
		leaves []*Tree
	}{
		{
			&Tree{
				"S", []*Tree{
					{"NP", []*Tree{{"NNP", []*Tree{p000}}}},
					{"VP", []*Tree{{"V", []*Tree{p100}}}},
					{".", []*Tree{p20}},
				},
			},
			[]*Tree{p000, p100, p20},
		},
	}

	for _, c := range cases {
		ls := c.tree.Leaves()

		if len(c.leaves) != len(ls) {
			t.Errorf("%v expected but got %v\n", c.leaves, ls)
		}

		for k, l := range c.leaves {
			if ls[k] != l {
				t.Errorf("%v expected but got %v\n", l, ls[k])
			}
		}
	}
}

func TestTree_LeafPositions(t *testing.T) {
	// TODO implement
}

func TestTree_Edges(t *testing.T) {
	// TODO implement
}

func TestTree_EdgePositions(t *testing.T) {
	// TODO implement
}

func TestTree_Sentence(t *testing.T) {
	// TODO implement
}

func TestTree_String(t *testing.T) {
	p := NewDecoder()

	cases := []struct {
		tree *Tree
		str  string
	}{
		{
			p.MustDecode("(S (NP (NNP John)) (VP (V runs)) (. .))"),
			"(S (NP (NNP John)) (VP (V runs)) (. .))",
		},
	}

	for _, c := range cases {
		s := c.tree.String()

		if c.str != s {
			t.Errorf("%v expected but got %v\n", c.str, s)
		}
	}
}

func BenchmarkTree_Leaves(b *testing.B) {
	p := NewDecoder()
	t := p.MustDecode("(S (NP (NNP John)) (VP (V runs)) (. .))")

	for i := 0; i < b.N; i++ {
		t.Leaves()
	}
}

func BenchmarkTree_LeafPositions(b *testing.B) {
	p := NewDecoder()
	t := p.MustDecode("(S (NP (NNP John)) (VP (V runs)) (. .))")

	for i := 0; i < b.N; i++ {
		t.LeafPositions()
	}
}
