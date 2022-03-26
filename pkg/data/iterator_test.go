package data

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"testing"
)

type sample struct {
	a, b string
}

var reader *csv.Reader
var factory func([]string) (interface{}, error)

func init() {
	file, err := os.Open("example.csv")

	if err != nil {
		log.Fatal(err)
	}

	reader = csv.NewReader(file)

	factory = func(record []string) (interface{}, error) {
		if len(record) != 2 {
			return nil, errors.New("unexpected record length")
		}

		return sample{
			record[0],
			record[1],
		}, nil
	}
}

func TestIterator_Next(t *testing.T) {
	i := NewIterator(reader, factory)

	if i.Next() && i.Sample() == nil {
		t.Errorf("unxpected empty sample")
	}
}

func TestIterator_Skip(t *testing.T) {
	i := NewIterator(reader, factory)

	n := i.Skip()
	err := i.Error()

	if n && err != nil {
		t.Errorf("unexected error %v", i.Error())
	}
}

func TestIterator_Sample(t *testing.T) {
	i := NewIterator(reader, factory)

	defer func() {
		recover()
	}()

	i.Sample()

	t.Errorf("should have panicked")
}
