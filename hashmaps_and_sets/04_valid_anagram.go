package hashmaps_and_sets

func isAnagram(s string, t string) bool {
	lenS := len(s)
	lenT := len(t)

	if lenS != lenT {
		return false
	}

	sM := make(map[rune]int, lenS)

	for _, inM := range s {
		sM[inM] = sM[inM] + 1
	}

	for _, tI := range t {
		if sM[tI] > 0 {
			sM[tI] = sM[tI] - 1
		} else {
			return false
		}
	}

	return true
}
