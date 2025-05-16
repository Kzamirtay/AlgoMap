package arrays_and_strings

func isSubsequence(s string, t string) bool {
	var wordStart int

	for _, sLet := range s {
		t = t[wordStart:]

		index, exists := letExists(sLet, t)
		if !exists {
			return false
		}

		wordStart = index + 1
	}
	return true
}

func letExists(let int32, word string) (int, bool) {
	for index, wLet := range word {
		if let == wLet {
			return index, true
		}
	}
	return 0, false
}
