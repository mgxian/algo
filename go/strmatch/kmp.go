package strmatch

// KMPSearch is a implement of Knuth Morris Pratt search
func KMPSearch(s, pattern string) int {
	next := getNext(pattern)
	j := 0
	for i := 0; i < len(s); i++ {
		for j > 0 && s[i] != pattern[j] {
			j = next[j-1] + 1
		}

		if s[i] == pattern[j] {
			j++
		}

		if j == len(pattern) {
			return i - (len(pattern) - 1)
		}
	}

	return -1
}

func getNext(pattern string) []int {
	next := make([]int, len(pattern))
	next[0] = -1
	k := -1
	for i := 1; i < len(pattern); i++ {
		for k != -1 && pattern[k+1] != pattern[i] {
			k = next[k]
		}

		if pattern[k+1] == pattern[i] {
			k++
		}

		next[i] = k
	}
	return next
}
