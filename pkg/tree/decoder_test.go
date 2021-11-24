package tree

import "testing"

func (d *decoder) MustDecode(s string) *Tree {
	if t, err := d.Decode(s); err != nil {
		panic(err)
	} else {
		return t
	}
}

func TestDecoder_Decode(t *testing.T) {
	cases := []struct {
		str  string
		tree *Tree
	}{
		{
			"(S (NP (NNP John)) (VP (V runs)) (. .))",
			&Tree{
				"S", []*Tree{
					{"NP", []*Tree{{"NNP", []*Tree{{"John", nil}}}}},
					{"VP", []*Tree{{"V", []*Tree{{"runs", nil}}}}},
					{".", []*Tree{{".", nil}}},
				},
			},
		},
	}

	d := NewDecoder()

	for _, c := range cases {
		tr := d.MustDecode("(S (NP (NNP John)) (VP (V runs)) (. .))")

		if !c.tree.Equals(tr) {
			t.Errorf("%v expected but got %v\n", c.tree, tr)
		}

		if tr.String() != c.tree.String() {
			t.Errorf("%v expected but got %v\n", c.tree, tr)
		}
	}
}

func TestParser_ParseFromStringError(t *testing.T) {
	cases := []struct {
		str string
	}{
		{
			"", // expected "(" got "end-of-string"
		},
		{
			"(", // expected ")" got "end-of-string"
		},
		{
			")", // expected "(" got ")"
		},
		{
			"John", // expected "(" got "foo"
		},
		{
			"(NNP John)(", // expected "end-of-string" got "("
		},
		{
			"(NNP John))", // expected "end-of-string" got "("
		},
	}

	p := NewDecoder()

	for _, c := range cases {
		tree, err := p.Decode(c.str)

		if err == nil {
			t.Errorf("error expected but got %v\n", tree)
		}
	}
}

func BenchmarkDecoder_ParseDecoder(b *testing.B) {
	d := NewDecoder()

	for i := 0; i < b.N; i++ {
		_, _ = d.Decode("(S (NP (NNP John)) (VP (V runs)) (. .))")
	}
}
