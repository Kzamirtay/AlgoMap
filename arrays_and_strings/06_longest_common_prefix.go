package arrays_and_strings

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	answer := &strings.Builder{}
	isset := false
	firstWord := strs[0]
	otherWords := strs[1:]

	for index, letter := range firstWord {
		for _, word := range otherWords {
			isset = false
			if len(word) > index && rune(word[index]) == letter {
				isset = true
			} else {
				break
			}
		}
		if isset {
			answer.WriteString(string(letter))
		} else {
			break
		}
	}

	return answer.String()
}
