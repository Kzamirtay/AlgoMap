package hashmaps_and_sets

func majorityElement(nums []int) int {
	var ans int
	nM := make(map[int]int, len(nums))

	for _, num := range nums {
		nM[num] += 1
	}

	for num, val := range nM {
		if nM[ans] < val {
			ans = num
		}
	}

	return ans
}
