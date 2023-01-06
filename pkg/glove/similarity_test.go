package glove

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

var glove map[string]*Embedding

func init() {
	file, err := os.Open("glove.6B.300d.txt")

	if err != nil {
		log.Fatal(err)
	}

	glove = make(map[string]*Embedding)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		embedding, err := NewEmbedding(strings.Split(scanner.Text(), " "), 300)

		if err != nil {
			log.Fatal(err)
		}

		glove[embedding.Token] = embedding
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
func TestSimilarity(t *testing.T) {
	s1 := strings.Split("king", " ")
	s2 := strings.Split("queen", " ")

	e1, _ := NewEmbedding(nil, 300)
	e2, _ := NewEmbedding(nil, 300)

	for _, t := range s1 {
		e1.Append(glove[t])
	}

	for _, t := range s2 {
		e2.Append(glove[t])
	}

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(Similarity(e1, e2))
}
