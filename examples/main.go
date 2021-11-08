package main

import (
	"fmt"
	"github.com/jonasknobloch/jinn/pkg/corenlp"
	"github.com/jonasknobloch/jinn/pkg/tree"
	"net/url"
)

func main() {
	u, _ := url.Parse("http://localhost:9000")

	c, _ := corenlp.NewClient(u, corenlp.Properties{
		Annotators:   corenlp.Annotators{corenlp.ParserAnnotator},
		OutputFormat: corenlp.FormatJSON,
	})

	d, err := c.Annotate("The quick brown fox jumped over the lazy dog.")

	if err != nil {
		panic(err)
	}

	p := d.Sentences[0].Parse

	dec := tree.NewDecoder()
	t, _ := dec.Decode(p)
	ht := tree.NewHybridTree(t, t.Leaves())

	fmt.Printf("%s\n\n%s\n\n%s", p, t.String(), ht.Sentence())
}
