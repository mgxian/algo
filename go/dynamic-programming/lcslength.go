package dynamic

func lcsLength(a, b string) int {
	n := len(a)
	m := len(b)
	maxLcs := make([][]int, n)
	for i := range maxLcs {
		maxLcs[i] = make([]int, m)
	}

	for j := 0; j < m; j++ {
		if a[0] == b[j] {
			maxLcs[0][j] = 1
		} else if j != 0 {
			maxLcs[0][j] = maxLcs[0][j-1]
		} else {
			maxLcs[0][j] = 0
		}
	}

	for i := 1; i < n; i++ {
		if a[i] == b[0] {
			maxLcs[i][0] = 1
		} else {
			maxLcs[i][0] = maxLcs[i-1][0]
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if a[i] == b[j] {
				maxLcs[i][j] = max(maxLcs[i-1][j], maxLcs[i][j-1], maxLcs[i-1][j-1]+1)
			} else {
				maxLcs[i][j] = max(maxLcs[i-1][j], maxLcs[i][j-1], maxLcs[i-1][j-1])
			}
		}
	}

	return maxLcs[n-1][m-1]
}

func max(a, b, c int) int {
	t := a
	if t < b {
		t = b
	}
	if t < c {
		t = c
	}
	return t
}
