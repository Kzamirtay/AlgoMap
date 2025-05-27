package hashmaps_and_sets

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	sM := make(map[string][]string)
	var key []rune
	var ans [][]string

	for _, str := range strs {
		key = []rune(str)

		slices.Sort(key)

		sM[string(key)] = append(sM[string(key)], str)
	}

	for _, strings := range sM {
		ans = append(ans, strings)
	}

	return ans
}
