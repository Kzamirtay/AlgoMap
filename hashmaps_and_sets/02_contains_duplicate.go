package hashmaps_and_sets

func containsDuplicate(nums []int) bool {
	nM := make(map[int]bool, len(nums))

	for _, n := range nums {
		if nM[n] {
			return true
		}
		nM[n] = true
	}

	return false
}
