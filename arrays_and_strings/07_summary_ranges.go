package arrays_and_strings

import "strconv"

func summaryRanges(nums []int) []string {
	var (
		answer          []string
		rangeStartIndex int
		rangeEndIndex   int
		rangeStop       bool
		nextIsset       bool
	)

	for index, num := range nums {
		rangeStop = false
		nextIsset = index < len(nums)-1

		if !nextIsset {
			rangeEndIndex = index
			rangeStop = true
		}

		// Если следующее число не равно (нынешнему числу + 1)
		if nextIsset && nums[index+1] != num+1 {
			rangeEndIndex = index
			rangeStop = true
		}

		if rangeStop {
			if nums[rangeStartIndex] == nums[rangeEndIndex] {
				answer = append(answer, strconv.Itoa(nums[rangeStartIndex]))
			} else {
				answer = append(answer, strconv.Itoa(nums[rangeStartIndex])+"->"+strconv.Itoa(nums[rangeEndIndex]))
			}

			if nextIsset {
				rangeStartIndex = index + 1
			} else {
				rangeStartIndex = rangeEndIndex
			}
		}
	}

	return answer
}
