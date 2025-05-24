package arrays_and_strings

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	res := &strings.Builder{}
	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		res.WriteString(string(word1[i]))
		res.WriteString(string(word2[j]))
		i++
		j++
	}

	for i < len(word1) {
		res.WriteString(string(word1[i]))
		i++
	}

	for j < len(word2) {
		res.WriteString(string(word2[j]))
		j++
	}

	return res.String()
}
