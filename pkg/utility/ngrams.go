package utility

func NumNGrams(e []string, n int) int {
	if n > len(e) || n == 0 {
		return 0
	}

	return len(e) - n + 1
}

func NGrams(e []string, n int, gs [][]string) [][]string {
	if n > len(e) {
		return [][]string{}
	}

	if gs == nil {
		gs = make([][]string, NumNGrams(e, n))
	}

	for i := 0; i < NumNGrams(e, n); i++ {
		gs[i] = e[i : i+n]
	}

	return gs
}
