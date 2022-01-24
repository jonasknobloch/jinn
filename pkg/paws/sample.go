package paws

import (
	"errors"
	"strconv"
)

type Sample struct {
	ID        int
	Sentence1 string
	Sentence2 string
	Label     bool
}

func NewSample(record []string) (Sample, error) {
	if len(record) != 4 {
		return Sample{}, errors.New("unexpected record length")
	}

	out := Sample{}

	if i, err := strconv.Atoi(record[0]); err != nil {
		return Sample{}, err
	} else {
		out.ID = i
	}

	out.Sentence1 = record[1]
	out.Sentence2 = record[2]

	if record[3] != "0" && record[3] != "1" {
		return Sample{}, errors.New("unexpected label")
	}

	out.Label = record[3] == "1"

	return out, nil
}
