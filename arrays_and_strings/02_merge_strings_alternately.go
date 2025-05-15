package arrays_and_strings

func mergeAlternately(word1 string, word2 string) string {
	answer := ""

	lenWord1 := len(word1)
	lenWord2 := len(word2)

	if lenWord1 != lenWord2 {
		var maxWord string
		var minWord string

		if lenWord1 > lenWord2 {
			maxWord = word1
			minWord = word2
		} else {
			maxWord = word2
			minWord = word1
		}

		for i, letter := range word1 {
			answer += string(letter)
			answer += string(word2[i])

			if len(minWord)-1 == i {
				break
			}
		}

		answer += maxWord[len(minWord):]
	} else {
		for i, letter := range word1 {
			answer += string(letter)
			answer += string(word2[i])
		}
	}

	return answer
}
