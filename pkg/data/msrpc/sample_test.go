package msrpc

import "testing"

func TestNewSample(t *testing.T) {
	record := []string{
		"1",
		"579975",
		"579810",
		"The DVD-CCA then appealed to the state Supreme Court.",
		"The DVD CCA appealed that decision to the U.S. Supreme Court.",
	}

	sample := Sample{
		Quality: true,
		ID1:     579975,
		ID2:     579810,
		String1: "The DVD-CCA then appealed to the state Supreme Court.",
		String2: "The DVD CCA appealed that decision to the U.S. Supreme Court.",
	}

	s, err := NewSample(record)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if *s.(*Sample) != sample {
		t.Errorf("struct contents do not match")
	}
}
