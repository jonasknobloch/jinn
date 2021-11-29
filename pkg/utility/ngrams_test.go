package utility

import (
	"reflect"
	"testing"
)

func TestNGrams(t *testing.T) {
	cases := []struct {
		input []string
		n     int
		gs    [][]string
	}{
		{
			[]string{"foo", "bar", "baz"},
			1,
			[][]string{{"foo"}, {"bar"}, {"baz"}},
		},
		{
			[]string{"foo", "bar", "baz"},
			2,
			[][]string{{"foo", "bar"}, {"bar", "baz"}},
		},
		{
			[]string{"foo", "bar", "baz"},
			3,
			[][]string{{"foo", "bar", "baz"}},
		},
		{
			[]string{"foo", "bar", "baz"},
			4,
			[][]string{},
		},
		{
			[]string{"foo", "bar", "baz"},
			0,
			[][]string{},
		},
	}

	for _, c := range cases {
		gs := NGrams(c.input, c.n, nil)

		if !reflect.DeepEqual(gs, c.gs) {
			t.Errorf("expected %v got %v", c.gs, gs)
		}
	}
}

func BenchmarkNGrams(b *testing.B) {
	e := []string{"foo", "bar", "baz"}

	for i := 0; i < b.N; i++ {
		_ = NGrams(e, 2, nil)
	}
}
