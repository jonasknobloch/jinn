package msrpc

import (
	"testing"
)

func TestNewIterator(t *testing.T) {
	_, err := NewIterator("msr_paraphrase_train.txt")

	if err != nil {
		t.Errorf("")
	}
}

func TestIterator_Next(t *testing.T) {
	i, err := NewIterator("msr_paraphrase_train.txt")

	if err != nil {
		panic(err)
	}

	if !i.Next() {
		t.Errorf("")
	}
}

func TestIterator_Sample(t *testing.T) {
	i, err := NewIterator("msr_paraphrase_train.txt")

	if err != nil {
		panic(err)
	}

	n := i.Next()
	e := i.Error()

	if !n && e == nil {
		panic("empty test data")
	}

	if n && e != nil {
		t.Errorf("unexpected error %s", e)
	}

	if n && i.Sample() == (Sample{}) {
		t.Errorf("unxpected empty sample")
	}
}

func TestIterator_SamplePanic(t *testing.T) {
	i, err := NewIterator("msr_paraphrase_train.txt")

	if err != nil {
		panic(err)
	}

	defer func() { recover() }()

	i.Sample()

	t.Errorf("should have panicked")
}

func TestIterator_Error(t *testing.T) {
	i, err := NewIterator("msr_paraphrase_train.txt")

	if err != nil {
		panic(err)
	}

	for i.Next() {
	}

	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
}
