package arrays_and_strings

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	var ans []int
	i, j := 0, 0

	UP, RIGHT, DOWN, LEFT := 0, 1, 2, 3

	direction := RIGHT

	UpWall := 0
	RightWall := n
	DownWall := m
	LeftWall := -1
	for len(ans) != m*n {
		if direction == RIGHT {
			for j < RightWall {
				ans = append(ans, matrix[i][j])
				j++
			}
			i, j = i+1, j-1
			RightWall--
			direction = DOWN
		} else if direction == DOWN {
			for i < DownWall {
				ans = append(ans, matrix[i][j])
				i++
			}
			i, j = i-1, j-1
			DownWall--
			direction = LEFT
		} else if direction == LEFT {
			for j > LeftWall {
				ans = append(ans, matrix[i][j])
				j--
			}
			i, j = i-1, j+1
			LeftWall++
			direction = UP
		} else {
			for i > UpWall {
				ans = append(ans, matrix[i][j])
				i--
			}
			i, j = i+1, j+1
			UpWall++
			direction = RIGHT
		}
	}

	return ans
}
