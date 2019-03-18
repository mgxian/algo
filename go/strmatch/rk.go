package strmatch

// RKSearch is a implement of Rabin-Karp search
func RKSearch(s, pattern string) int {
	if len(s) < len(pattern) {
		return -1
	}

	const CharacterSize int = 26
	patternLength := len(pattern)
	patternPow := make([]int, patternLength)
	patternPow[0] = 1
	for i := 1; i < patternLength; i++ {
		patternPow[i] = CharacterSize * patternPow[i-1]
	}

	patternHash := 0
	for i, v := range pattern {
		idx := patternLength - 1 - i
		patternHash += patternPow[idx] * (int(v) - 'a')
	}

	preHash := 0
	for i := 0; i < patternLength; i++ {
		idx := patternLength - 1 - i
		preHash += patternPow[idx] * (int(s[i]) - 'a')
	}

	if preHash == patternHash {
		return 0
	}

	for i := 1; i <= len(s)-patternLength; i++ {
		hash := 26*(preHash-patternPow[patternLength-1]*(int(s[i-1])-'a')) + patternPow[0]*(int(s[i+patternLength-1])-'a')
		if hash == patternHash {
			return i
		}
		preHash = hash
	}

	return -1
}
