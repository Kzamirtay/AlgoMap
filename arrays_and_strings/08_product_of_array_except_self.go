package arrays_and_strings

import "slices"

func productExceptSelf(nums []int) []int {
	numsLen := len(nums)
	numsLastI := numsLen - 1

	left := make([]int, 0, numsLen)
	left = append(left, 1)

	right := make([]int, 0, numsLen)
	right = append(right, 1)

	for i := 0; i < len(nums)-1; i++ {
		left = append(left, left[i]*nums[i])
		right = append(right, right[i]*nums[numsLastI-i])
	}

	// +
	slices.Reverse(right)

	for i := 0; i < len(nums); i++ {
		nums[i] = left[i] * right[i]
	}

	return nums
}
