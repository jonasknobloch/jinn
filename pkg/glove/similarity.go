package glove

import "math"

func Scalar(a, b *Embedding) float64 {
	if len(a.Vector) != len(b.Vector) {
		panic("mismatched vector dimensions")
	}

	var sum float64

	for i := 0; i < len(a.Vector); i++ {
		sum += a.Vector[i] * b.Vector[i]
	}

	return sum
}

func Norm(a *Embedding) float64 {
	var sum float64

	for _, v := range a.Vector {
		sum += v * v
	}

	return math.Sqrt(sum)
}

func Similarity(a, b *Embedding) float64 {
	return Scalar(a, b) / (Norm(a) * Norm(b))
}
