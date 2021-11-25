package msrpc

import (
	"errors"
	"strconv"
	"strings"
)

type Sample struct {
	Quality bool
	ID1     int
	ID2     int
	String1 string
	String2 string
}

func NewSample(line string) (Sample, error) {
	data := strings.Split(line, "\t")
	out := Sample{}

	if data[0] != "0" && data[0] != "1" {
		return Sample{}, errors.New("unknown quality")
	}

	out.Quality = data[0] == "1"

	if id, err := strconv.Atoi(data[1]); err == nil {
		out.ID1 = id
	} else {
		return Sample{}, err
	}

	if id, err := strconv.Atoi(data[2]); err == nil {
		out.ID2 = id
	} else {
		return Sample{}, err
	}

	out.String1 = data[3]
	out.String2 = data[4]

	return out, nil
}
