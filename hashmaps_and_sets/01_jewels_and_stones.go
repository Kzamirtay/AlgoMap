package hashmaps_and_sets

func numJewelsInStones(jewels string, stones string) int {
	var answer int
	jM := make(map[rune]bool, len(jewels))

	for _, j := range jewels {
		jM[j] = true
	}

	for _, s := range stones {
		if jM[s] {
			answer++
		}
	}

	return answer
}
