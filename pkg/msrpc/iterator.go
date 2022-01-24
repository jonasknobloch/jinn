package msrpc

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type Iterator struct {
	scanner *bufio.Scanner
	sample  Sample
	error   error
}

func NewIterator(name string) (*Iterator, error) {
	file, err := os.Open(name)

	if err != nil {
		return nil, err
	}

	i := &Iterator{
		scanner: bufio.NewScanner(file),
		sample:  Sample{},
	}

	i.scanner.Scan()

	if err := i.scanner.Err(); err != nil {
		return nil, err
	}

	header := strings.TrimLeft(i.scanner.Text(), "\uFEFF")

	if header != "Quality\t#1 ID\t#2 ID\t#1 String\t#2 String" {
		return nil, errors.New("unknown file contents")
	}

	return i, nil
}

func (i *Iterator) Next() bool {
	if !i.scanner.Scan() {
		i.error = i.scanner.Err()

		return false
	}

	i.sample, i.error = NewSample(i.scanner.Text())

	if i.error != nil {
		return false
	}

	return true
}

func (i *Iterator) Error() error {
	return i.error
}

func (i *Iterator) Sample() Sample {
	if i.sample == (Sample{}) {
		panic("Sample called before Next")
	}

	return i.sample
}
