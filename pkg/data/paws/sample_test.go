package paws

import "testing"

func TestNewSample(t *testing.T) {
	record := []string{
		"5",
		"It is the seat of Zerendi District in Akmola Region .",
		"It is the seat of the district of Zerendi in Akmola region .",
		"1",
	}

	sample := &Sample{
		ID:        5,
		Sentence1: "It is the seat of Zerendi District in Akmola Region .",
		Sentence2: "It is the seat of the district of Zerendi in Akmola region .",
		Label:     true,
	}

	s, err := NewSample(record)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if *s.(*Sample) != *sample {
		t.Fatalf("struct contents do not match")
	}
}
