package hashmaps_and_sets

func isValidSudoku(board [][]byte) bool {
	rM := make(map[[2]int]int, 10)
	cM := make(map[[2]int]int, 10)
	sM := make(map[[3]int]int, 10)

	key1 := [2]int{}
	key2 := [3]int{}

	var elem byte

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			elem = board[i][j]
			if elem != '.' {

				key1 = [2]int{i, int(elem)}
				rM[key1] = rM[key1] + 1
				if rM[key1] > 1 {
					return false
				}

				key1 = [2]int{j, int(elem)}
				cM[key1] = cM[key1] + 1
				if cM[key1] > 1 {
					return false
				}

				key2 = [3]int{i / 3, j / 3, int(elem)}
				sM[key2] = sM[key2] + 1
				if sM[key2] > 1 {
					return false
				}
			}
		}
	}

	return true
}
