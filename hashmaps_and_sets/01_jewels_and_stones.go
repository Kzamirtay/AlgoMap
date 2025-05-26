package hashmaps_and_sets

func numJewelsInStones(jewels string, stones string) int {
	var answer int

	for _, stone := range stones {
		for _, jewel := range jewels {
			if stone == jewel {
				answer++
			}
		}
	}

	return answer
}
