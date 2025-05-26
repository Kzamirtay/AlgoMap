package hashmaps_and_sets

func twoSum(nums []int, target int) []int {
	nM := make(map[int]int, len(nums))

	for i, n := range nums {
		nM[n] = i
	}

	for i, n := range nums {
		if nM[target-n] > 0 && i != nM[target-n] {
			return []int{i, nM[target-n]}
		}
	}

	return []int{}
}
