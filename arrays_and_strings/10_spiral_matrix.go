package arrays_and_strings

func spiralOrder(matrix [][]int) []int {
	var answer []int
	rowI, colI := 0, 0
	up := 0
	right := len(matrix[0])
	down := len(matrix)
	left := 0

	for {
		answer = append(answer, matrix[rowI][colI])
		colI++
		if rowI == up && colI == right {
			// Убираем Верх
			up++
		}
		if rowI == down && colI == right {
			// Убираем Право
			right--
		}
		if rowI == down && colI == left {
			// Убираем Низ
			down--
		}
		if rowI == up && colI == left {
			// Убираем Низ
			left++
		}
	}
	return answer
}
