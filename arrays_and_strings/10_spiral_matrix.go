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
		if rowI == up {
			if colI == right-1 {
				// Убираем Верх
				up++
				rowI++
			} else {
				colI++
			}
		} else if colI == right-1 {
			if rowI == down-1 {
				// Убираем Право
				right--
				colI--
			} else {
				rowI++
			}
		} else if rowI == down {
			if colI == left {
				// Убираем Низ
				down--
			} else {
				colI--
			}
		}
		//if colI == left {
		//	if rowI == up {
		//		// Убираем Низ
		//		left++
		//	} else {
		//		rowI++
		//	}
		//}
	}
	return answer
}
