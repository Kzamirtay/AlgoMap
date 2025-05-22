package arrays_and_strings

func spiralOrder(matrix [][]int) []int {
	var answer []int
	rowLen := len(matrix)
	colLen := len(matrix[0])
	rowLastI := rowLen - 1
	colLastI := colLen - 1
	rowI, colI := 0, 0
	rmCol, rmRow := true, true

	for rowI <= rowLen && colI <= colLen {
		answer = append(answer, matrix[rowI][colI])
		if rowLastI == rowI {
			if rmCol {
				colLastI--
				rmCol = false
				rmRow = true
			}
			colI--
		}
		if colLastI == colI {
			if rmRow {
				rowLastI--
				rmRow = false
				rmCol = true
			}
			rowI++
		}
		if colLastI > colI {
			colI++
		}
	}
	return answer
}
