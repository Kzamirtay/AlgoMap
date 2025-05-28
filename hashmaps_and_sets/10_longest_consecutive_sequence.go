package hashmaps_and_sets

func longestConsecutive(nums []int) int {
	nM := make(map[int]bool)

	for _, num := range nums {
		nM[num] = true
	}

	ans := 0
	lon := 1
	key := 0

	for num := range nM {
		if !nM[num-1] {
			key = num + 1
			for nM[key] {
				key = key + 1
				lon++
			}
		}
		if lon > ans {
			ans = lon
		}
		lon = 1
	}

	return ans
}
