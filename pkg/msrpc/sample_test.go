package msrpc

import (
	"reflect"
	"testing"
)

func TestNewSample(t *testing.T) {
	cases := []struct {
		input     string
		shouldErr bool
		sample    Sample
	}{
		{
			"1\t579975\t579810\tFoo\tBar\n",
			false,
			Sample{
				Quality: true,
				ID1:     579975,
				ID2:     579810,
				String1: "Foo",
				String2: "Bar",
			},
		},
		{
			"2\t579975\t579810\tFoo\tBar\n",
			true,
			Sample{},
		},
		{
			"1\t5799ab\t579810\tFoo\tBar\n",
			true,
			Sample{},
		},
		{
			"1\t579975\t5798cd\tFoo\tBar\n",
			true,
			Sample{},
		},
	}

	for _, c := range cases {
		s, err := NewSample(c.input)

		if c.shouldErr && err == nil {
			t.Errorf("expected error but got nil")
		}

		if !c.shouldErr && err != nil {
			t.Errorf("unexpected error %s", err)
		}

		if !c.shouldErr && err == nil {
			if reflect.DeepEqual(c.sample, s) {
				t.Errorf("expected %v but got %v", c.sample, s)
			}
		}
	}
}
