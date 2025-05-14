package arrays_and_strings

func findClosestNumber(nums []int) int {
	numsLen := len(nums) - 1
	answer := nums[0]

	for index, num := range nums {
		if index != numsLen && abs(answer) > abs(nums[index+1]) {
			answer = nums[index+1]
		}

		if abs(num) == abs(answer) {
			if num > answer {
				answer = num
			}
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
