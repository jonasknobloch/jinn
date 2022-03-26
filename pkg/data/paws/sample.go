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

// id
// sentence1
// sentence2
// label

func NewSample(record []string) (interface{}, error) {
	if len(record) != 4 {
		return nil, errors.New("unexpected record length")
	}

	sample := &Sample{}

	if i, err := strconv.Atoi(record[0]); err != nil {
		return nil, err
	} else {
		sample.ID = i
	}

	sample.Sentence1 = record[1]
	sample.Sentence2 = record[2]

	if record[3] != "0" && record[3] != "1" {
		return nil, errors.New("unexpected label")
	}

	sample.Label = record[3] == "1"

	return sample, nil
}
