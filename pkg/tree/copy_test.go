package tree

import "testing"

func TestTree_Copy(t *testing.T) {
	d := NewDecoder()

	cases := []*Tree{
		d.MustDecode("(S (NP (NNP John)) (VP (V runs)) (. .))"),
	}

	for _, tr := range cases {
		m1 := make(map[*Tree]*Tree)
		m2 := make(map[*Tree]*Tree)

		cpy := tr.Copy(m1)

		for k, v := range m1 {
			m2[v] = k
		}

		if !tr.Equals(cpy) {
			t.Errorf("%v expected got %v", tr, cpy)
		}

		cpy.Walk(func(st *Tree) {
			if _, ok := m1[st]; !ok {
				t.Errorf("no mapping for subtree %v", st)
			}
		})

		tr.Walk(func(st *Tree) {
			if _, ok := m2[st]; !ok {
				t.Errorf("no mapping for subtree %v", st)
			}
		})

		cpy.Walk(func(st *Tree) {
			if !st.Equals(m1[st]) {
				t.Errorf("%v expected but got %v", st, m1[st])
			}
		})

		tr.Walk(func(st *Tree) {
			if !st.Equals(m2[st]) {
				t.Errorf("%v expected but got %v", st, m2[st])
			}
		})
	}
}
