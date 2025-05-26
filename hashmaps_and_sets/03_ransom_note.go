package hashmaps_and_sets

func canConstruct(ransomNote string, magazine string) bool {
	mp := make(map[rune]int, len(ransomNote))

	for _, inM := range magazine {
		mp[inM] = mp[inM] + 1
	}

	for _, inR := range ransomNote {
		if mp[inR] > 0 {
			mp[inR] = mp[inR] - 1
		} else {
			return false
		}
	}

	return true
}
