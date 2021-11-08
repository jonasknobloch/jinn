package tree

import "strings"

type HybridTree struct {
	tree  *Tree
	order []*Tree
}

func NewHybridTree(t *Tree, o []*Tree) *HybridTree {
	return &HybridTree{
		tree:  t,
		order: o,
	}
}

func (ht *HybridTree) Sentence() string {
	var sb strings.Builder

	for _, n := range ht.Cover(nil) {
		sb.WriteString(" ")
		sb.WriteString(n.Label)
	}

	return sb.String()[1:]
}

func (ht *HybridTree) Cover(p []int) []*Tree {
	cover := make([]*Tree, 0)
	leaves := ht.tree.SubtreeAtPosition(p).Leaves()

	for _, o := range ht.order {
		for _, l := range leaves {
			if o == l {
				cover = append(cover, o)
			}
		}
	}

	return cover
}
