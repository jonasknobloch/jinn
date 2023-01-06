package glove

import (
	"strconv"
)

type Embedding struct {
	Token  string
	Vector []float64
}

func NewEmbedding(record []string, dimension int) (*Embedding, error) {
	embedding := &Embedding{
		Vector: make([]float64, dimension),
	}

	if len(record) == 0 {
		for i := range embedding.Vector {
			embedding.Vector[i] = 1
		}

		return embedding, nil
	}

	embedding.Token = record[0]

	for i, v := range record[1:] {
		f, err := strconv.ParseFloat(v, 64)

		if err != nil {
			return nil, err
		}

		embedding.Vector[i] = f
	}

	return embedding, nil
}

func (a *Embedding) Clone() *Embedding {
	b := &Embedding{
		Vector: make([]float64, len(a.Vector)),
	}

	b.Token = a.Token

	copy(b.Vector, a.Vector)

	return b
}

func (a *Embedding) Append(b *Embedding) {
	a.Token += " " + b.Token

	if len(a.Vector) != len(b.Vector) {
		panic("mismatched vector dimensions")
	}

	for i := range a.Vector {
		a.Vector[i] *= b.Vector[i]
	}
}
