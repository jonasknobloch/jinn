package paws_test

import (
	"github.com/jonasknobloch/jinn/pkg/paws"
	"log"
)

func ExampleIterator() {
	i, err := paws.NewIterator("train.tsv")

	if err != nil {
		log.Fatal(err)
	}

	for i.Next() {
		// do nothing
	}

	if err := i.Error(); err == nil {
		log.Fatal(err)
	}

	// Output:
}
