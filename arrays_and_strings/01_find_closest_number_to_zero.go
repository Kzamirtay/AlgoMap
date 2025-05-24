package arrays_and_strings

func findClosestNumber(nums []int) int {
	answer := nums[0]

	for index, num := range nums {
		if index != len(nums)-1 && abs(answer) > abs(nums[index+1]) {
			answer = nums[index+1]
		} else if abs(num) == abs(answer) && num > answer {
			answer = num
		}
	}

	return answer
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}
