package score

import (
	"math"
	"testing"
)

func TestModifiedPrecisionScore(t *testing.T) {
	cases := []struct {
		e  []string
		rs [][]string
		n  int
		p  float64
	}{
		{
			[]string{"foo", "bar"},
			[][]string{{"foo", "bar", "baz"}},
			1,
			1.0,
		},
		{
			[]string{"foo", "baz"},
			[][]string{{"foo", "bar", "baz"}},
			2,
			0,
		},
		{
			[]string{"foo", "bar"},
			[][]string{{"foo"}, {"bar"}},
			1,
			1.0,
		},
		{
			[]string{"foo", "bar"},
			[][]string{{"foo"}, {"bar"}},
			2,
			0,
		},
		{
			[]string{"foo", "bar"},
			[][]string{},
			1,
			0,
		},
	}

	for _, c := range cases {
		p := modifiedPrecisionScore(c.n, c.e, c.rs...)

		if p != c.p {
			t.Errorf("expected %f got %f", c.p, p)
		}
	}
}

func TestCombinedModifiedPrecisionScore(t *testing.T) {
	cases := []struct {
		e  []string
		rs [][]string
		p  float64
	}{
		{
			[]string{"foo", "bar", "foo"},
			[][]string{{"foo"}, {"foo", "bar"}, {"foo", "bar", "baz"}},
			0,
		},
		{
			[]string{"foo", "bar", "foo", "baz"},
			[][]string{{"foo", "bar"}, {"foo", "bar", "foo", "baz"}},
			1,
		},
		{
			[]string{"foo", "bar", "baz", "foo", "bar", "baz"},
			[][]string{{"foo", "bar", "baz", "foo"}, {"baz", "foo", "bar", "baz"}},
			math.Pow(5/6.0*3/5.0*3/4.0*2/3.0, 1/4.0),
		},
		{
			[]string{"foo", "bar", "foo"},
			[][]string{},
			0,
		},
	}

	for _, c := range cases {
		p := combinedModifiedPrecisionScore(c.e, c.rs...)

		if p != c.p {
			t.Errorf("expected %f got %f", c.p, p)
		}
	}
}

func TestEffectiveReferencesLength(t *testing.T) {
	cases := []struct {
		e   []string
		rs  [][]string
		erl int
	}{
		{
			[]string{"foo", "bar"},
			[][]string{{"foo", "bar"}},
			2,
		},
		{
			[]string{"foo", "bar"},
			[][]string{{"foo"}, {"foo", "bar", "baz"}},
			1,
		},
		{
			[]string{"foo", "bar"},
			[][]string{},
			0,
		},
	}

	for _, c := range cases {
		erl := effectiveReferencesLength(c.e, c.rs...)

		if erl != c.erl {
			t.Errorf("expected %d got %d", c.erl, erl)
		}
	}
}

func TestBrevityPenalty(t *testing.T) {
	cases := []struct {
		e  []string
		rs [][]string
		bp float64
	}{
		{
			[]string{"foo", "bar", "baz"},
			[][]string{{"foo", "bar"}},
			1,
		},
		{
			[]string{"foo", "bar", "baz"},
			[][]string{{"foo", "bar", "baz"}},
			math.Exp(1 - (3 / 3.0)),
		},
		{
			[]string{"foo", "bar"},
			[][]string{{"foo", "bar", "baz"}},
			math.Exp(1 - (3 / 2.0)),
		},
		{
			[]string{"foo", "bar"},
			[][]string{},
			1,
		},
	}

	for _, c := range cases {
		bp := brevityPenalty(c.e, c.rs...)

		if bp != c.bp {
			t.Errorf("expected %f got %f", c.bp, bp)
		}
	}
}

func TestBLEU(t *testing.T) {
	cases := []struct {
		e    []string
		rs   [][]string
		bleu float64
	}{
		{
			[]string{"foo", "bar", "baz"},
			[][]string{{"foo", "bar", "baz"}},
			1,
		},
		{
			[]string{"foo", "bar", "baz", "foo", "bar", "baz"},
			[][]string{{"foo", "bar", "baz", "foo"}, {"baz", "foo", "bar", "baz"}},
			1 * math.Pow(1/4.0, 1/4.0),
		},
		{
			[]string{"foo", "bar", "baz"},
			[][]string{},
			0,
		},
	}

	for _, c := range cases {
		bleu := BLEU(c.e, c.rs...)

		if bleu != c.bleu {
			t.Errorf("expected %f got %f", c.bleu, bleu)
		}
	}
}
