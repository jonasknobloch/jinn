package data

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func Example() {
	file, err := os.Open("example.csv")

	if err != nil {
		log.Fatal(err)
	}

	iterator := NewIterator(csv.NewReader(file), factory)

	for iterator.Next() {
		s := iterator.Sample().(sample)

		fmt.Printf("%s %s\n", s.a, s.b)
	}

	// Output:
	// foo bar
	// bar baz
}
