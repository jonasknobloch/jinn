package score

import (
	"github.com/jonasknobloch/jinn/pkg/utility"
	"math"
	"strings"
)

func modifiedPrecisionScore(n int, e []string, rs ...[]string) float64 {
	count := func(nGrams [][]string) map[string]int {
		o := make(map[string]int)

		for _, n := range nGrams {
			k := strings.Join(n, " ")

			if _, ok := o[k]; ok {
				o[k]++
			} else {
				o[k] = 1
			}
		}

		return o
	}

	max := func(rs ...map[string]int) map[string]int {
		m := make(map[string]int)

		for _, r := range rs {
			for k, v1 := range r {
				if v2, ok := m[k]; !ok || v1 > v2 {
					m[k] = v1
				}
			}
		}

		return m
	}

	intersect := func(a, b map[string]int) map[string]int {
		i := make(map[string]int)

		for k, v1 := range a {
			if v2, ok := b[k]; ok {
				if v1 < v2 {
					i[k] = v1
				} else {
					i[k] = v2
				}
			}
		}

		return i
	}

	sum := func(o map[string]int) int {
		s := 0

		for _, v := range o {
			s += v
		}

		return s
	}

	Rs := make([]map[string]int, len(rs))

	for k, r := range rs {
		Rs[k] = count(utility.NGrams(r, n, nil))
	}

	E := count(utility.NGrams(e, n, nil))

	if len(E) == 0 {
		return 0
	}

	return float64(sum(intersect(E, max(Rs...)))) / float64(sum(E))
}

func combinedModifiedPrecisionScore(e []string, rs ...[]string) float64 {
	n := 4

	if n > len(e) {
		n = len(e)
	}

	var sum float64

	for i := 1; i <= n; i++ {
		mps := modifiedPrecisionScore(i, e, rs...)

		if mps == 0 {
			return 0
		}

		sum += math.Log(mps)
	}

	return math.Exp(sum / float64(n))
}

func effectiveReferencesLength(e []string, rs ...[]string) int {
	abs := func(a, b int) int {
		if a < b {
			return a
		}

		return b
	}

	var minR, minL int

	for _, r := range rs {
		if len(r) > minR {
			minR = len(r)
			minL = len(r)
		}
	}

	for _, r := range rs {
		if len(r) > minR {
			break
		}

		minR = len(r)
		a := abs(len(e), len(r))

		if a > minL {
			break
		}

		minL = a
	}

	return minR
}

func brevityPenalty(e []string, rs ...[]string) float64 {
	erl := effectiveReferencesLength(e, rs...)

	if len(e) > erl {
		return 1
	}

	return math.Exp(float64(1) - float64(erl)/float64(len(e)))
}

func BLEU(e []string, r ...[]string) float64 {
	return brevityPenalty(e, r...) * combinedModifiedPrecisionScore(e, r...)
}
