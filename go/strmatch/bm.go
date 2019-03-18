package strmatch

// BMSearch is a implement of Boyer-Moore search
func BMSearch(s, pattern string) int {

	return -1
}

func generateBC(pattern string) []int {
	const CharacterSize int = 256
	bc := make([]int, CharacterSize)

	for i := range pattern {
		bc[i] = -1
	}

	for i, v := range pattern {
		bc[int(v)] = i
	}

	return bc
}

func generateGS(pattern string) ([]int, []bool) {
	patternLength := len(pattern)
	suffix := make([]int, patternLength)
	prefix := make([]bool, patternLength)

	for i := range suffix {
		suffix[i] = -1
	}

	for i := 0; i < patternLength-1; i++ {
		j := i
		k := 0
		for j >= 0 && pattern[j] == pattern[patternLength-1-k] {
			j--
			k++
			suffix[k] = j + 1
		}

		if j == -1 {
			prefix[k] = true
		}
	}

	return suffix, prefix
}
