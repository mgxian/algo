package dynamic

func levenshteinDistance(a, b string) int {
	n := len(a)
	m := len(b)
	minDist := make([][]int, n)
	for i := range minDist {
		minDist[i] = make([]int, m)
	}

	for j := 0; j < m; j++ {
		if j == 0 && a[0] == b[j] {
			minDist[0][j] = 0
		} else if j == 0 && a[0] != b[j] {
			minDist[0][j] = 1
		} else {
			minDist[0][j] = minDist[0][j-1] + 1
		}
	}

	for i := 1; i < n; i++ {
		minDist[i][0] = minDist[i-1][0] + 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if a[i] == b[j] {
				minDist[i][j] = min(minDist[i-1][j]+1, minDist[i][j-1]+1, minDist[i-1][j-1])
			} else {
				minDist[i][j] = min(minDist[i-1][j]+1, minDist[i][j-1]+1, minDist[i-1][j-1]+1)
			}
		}
	}

	return minDist[n-1][m-1]
}

func min(a, b, c int) int {
	t := a
	if t > b {
		t = b
	}
	if t > c {
		t = c
	}
	return t
}
