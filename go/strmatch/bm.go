package strmatch

// BMSearch is a implement of Boyer-Moore search
func BMSearch(s, pattern string) int {
	bc := generateBC(pattern)
	suffix, prefix := generateGS(pattern)
	i := 0
	for i <= len(s)-len(pattern) {
		j := len(pattern) - 1
		for j >= 0 {
			if s[i+j] != pattern[j] {
				break
			}
			j--
		}
		if j < 0 {
			return i
		}

		x := j - bc[s[i+j]]
		y := 0
		if j < len(pattern)-1 {
			y = moveByGS(j, len(pattern), suffix, prefix)
		}

		t := x
		if y > x {
			t = y
		}

		i += t
	}

	return -1
}

func generateBC(pattern string) []int {
	const CharacterSize int = 256
	bc := make([]int, CharacterSize)

	for i := range bc {
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
		prefix[i] = false
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

func moveByGS(j, patternLength int, suffix []int, prefix []bool) int {
	k := patternLength - 1 - j
	if suffix[k] != -1 {
		return j + 1 - suffix[k]
	}

	for i := j + 2; i <= patternLength-1; i++ {
		if prefix[patternLength-i] == true {
			return i
		}
	}

	return patternLength
}
