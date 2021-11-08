package tree

import "testing"

func TestPosition_String(t *testing.T) {
	cases := []struct {
		pos Position
		str string
	}{
		{
			Position{},
			"",
		},
		{
			Position{0},
			"0",
		},
		{
			Position{1, 2, 3},
			"123",
		},
	}

	for _, c := range cases {
		s := c.pos.String()

		if s != c.str {
			t.Errorf("%s expected but got %s\n", c.str, s)
		}
	}
}

func TestPositionFromString(t *testing.T) {
	// TODO implement
}
