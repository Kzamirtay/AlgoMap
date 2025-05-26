package hashmaps_and_sets

import "slices"

func maxNumberOfBalloons(text string) int {
	word := map[rune]int{
		'b': 1,
		'a': 1,
		'l': 2,
		'o': 2,
		'n': 1,
	}

	tM := make(map[rune]int, len(text))

	for _, let := range text {
		tM[let] = tM[let] + 1
	}

	counts := make([]int, 5)

	i := 0
	for let, size := range word {
		counts[i] = tM[let] / size
		i++
	}

	return slices.Min(counts)
}
