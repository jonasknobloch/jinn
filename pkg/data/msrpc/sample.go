package msrpc

import (
	"errors"
	"strconv"
)

type Sample struct {
	Quality bool
	ID1     int
	ID2     int
	String1 string
	String2 string
}

// Quality
// #1 ID
// #2 ID
// #1 String
// #2 String

func NewSample(record []string) (interface{}, error) {
	if len(record) != 5 {
		return nil, errors.New("unexpected record length")
	}

	sample := &Sample{}

	if record[0] != "0" && record[0] != "1" {
		return nil, errors.New("unknown quality")
	}

	sample.Quality = record[0] == "1"

	if id, err := strconv.Atoi(record[1]); err == nil {
		sample.ID1 = id
	} else {
		return nil, err
	}

	if id, err := strconv.Atoi(record[2]); err == nil {
		sample.ID2 = id
	} else {
		return nil, err
	}

	sample.String1 = record[3]
	sample.String2 = record[4]

	return sample, nil
}
