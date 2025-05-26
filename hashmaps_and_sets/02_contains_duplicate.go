package hashmaps_and_sets

import "sort"

func containsDuplicate(nums []int) bool {
	var lst = len(nums) - 1

	sort.Ints(nums)

	for i, num := range nums {
		if i != lst && num == nums[i+1] {
			return true
		}
	}

	return false
}
