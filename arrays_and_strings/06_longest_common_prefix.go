package arrays_and_strings

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	answer := ""
	isset := false
	firstWord := strs[0]
	otherWords := strs[1:]

	for index, letter := range firstWord {
		for _, word := range otherWords {
			isset = false

			if len(word) > index && int32(word[index]) == letter {
				isset = true
			} else {
				break
			}
		}
		if isset {
			answer += string(letter)
		} else {
			break
		}
	}

	return answer
}
