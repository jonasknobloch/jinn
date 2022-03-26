package data

import (
	"encoding/csv"
)

type Iterator struct {
	reader  *csv.Reader
	error   error
	sample  interface{}
	factory func([]string) (interface{}, error)
}

func NewIterator(reader *csv.Reader, factory func([]string) (interface{}, error)) *Iterator {
	return &Iterator{
		reader:  reader,
		sample:  nil,
		factory: factory,
	}
}

func (i *Iterator) Next() bool {
	record, err := i.reader.Read()

	if err != nil {
		i.error = err

		return false
	}

	i.sample, i.error = i.factory(record)

	if i.error != nil {
		return false
	}

	return true
}

func (i *Iterator) Skip() bool {
	_, err := i.reader.Read()

	if err != nil {
		i.error = err

		return false
	}

	return true
}

func (i *Iterator) Error() error {
	return i.error
}

func (i *Iterator) Sample() interface{} {
	if i.sample == nil {
		panic("Sample called before Next")
	}

	return i.sample
}
