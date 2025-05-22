package arrays_and_strings

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 2 <= intervals.length <= 104
	// intervals[i].length == 2
	// Слайсы отсортированы по first

	var answer [][]int
	answer = append(answer, intervals[0])

	var answerLastI int

	var back []int
	var current []int
	//Начинаем со второго индекса
	for i := 1; i < len(intervals); i++ {
		answerLastI = len(answer) - 1
		back = answer[answerLastI]

		current = intervals[i]

		if current[0] <= back[1] {
			if back[1] < current[1] {
				back = []int{back[0], current[1]}
			}
			answer[answerLastI] = back
		} else {
			answer = append(answer, current)
		}
	}

	return answer
}
