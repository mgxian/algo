package strmatch

func contain(main, pattern string) int {
	m := len(main)
	n := len(pattern)

	if m == 0 || n == 0 {
		return -1
	}

	for i := 0; i <= m-n; i++ {
		j := 0
		for j < n {
			if main[i+j] != pattern[j] {
				break
			}
			j++
		}
		if j == n {
			return i
		}
	}
	return -1
}
