package tree

import "testing"

func TestTree_Equals(t *testing.T) {
	d := NewDecoder()

	cases := []struct {
		t1     *Tree
		t2     *Tree
		equals bool
	}{
		{
			d.MustDecode("(S (NP (NNP John)) (VP (V runs)) (. .))"),
			d.MustDecode("(S (NP (NNP John)) (VP (V runs)) (. .))"),
			true,
		},
		{
			d.MustDecode("(S (NP (NNP John)) (VP (V runs)) (. .))"),
			d.MustDecode("(S (NP John) (VP runs) (. .))"),
			false,
		},
	}

	for _, c := range cases {
		e := c.t1.Equals(c.t2)

		if e != c.equals {
			t.Errorf("%t expected but got %t", e, c.equals)
		}
	}
}
