package arrays_and_strings

func romanToInt(s string) int {
	cat := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	answer := 0
	lenS := len(s)
	current := 0
	next := 0
	nextUsed := false

	for i, letter := range s {
		next = 0
		if nextUsed {
			nextUsed = false
			continue
		}

		current = cat[string(letter)]

		if lenS-1 != i {
			next = cat[string(s[i+1])]
		}

		if current < next {
			answer += next - current
			nextUsed = true
			continue
		}
		answer += current
	}

	return answer
}
