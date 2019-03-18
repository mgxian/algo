package strmatch

// BFSearch is brute force string search
func BFSearch(s, pattern string) int {
	for i := 0; i <= len(s)-len(pattern); i++ {
		j := 0
		for ; j < len(pattern); j++ {
			if s[i+j] != pattern[j] {
				break
			}
		}
		if j == len(pattern) {
			return i
		}
	}
	return -1
}
