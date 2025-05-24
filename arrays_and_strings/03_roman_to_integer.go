package arrays_and_strings

func romanToInt(s string) int {
	cat := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	str := []rune(s)
	lastI := len(s) - 1
	answer := cat[str[lastI]]

	for i := lastI; i > 0; i-- {
		if cat[str[i]] > cat[str[i-1]] {
			answer -= cat[str[i-1]]
		} else {
			answer += cat[str[i-1]]
		}
	}

	return answer
}
