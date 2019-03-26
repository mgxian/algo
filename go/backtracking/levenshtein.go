package backtracking

func levenshteinDistance(i, j, dist int, a, b string, min *int) {
	n := len(a)
	m := len(b)
	if i == n || j == m {
		if i < n {
			dist += n - i
		}
		if j < m {
			dist += m - j
		}
		if dist < *min {
			*min = dist
		}
		return
	}

	if a[i] == b[j] {
		levenshteinDistance(i+1, j+1, dist, a, b, min)
	} else {
		levenshteinDistance(i+1, j, dist+1, a, b, min)
		levenshteinDistance(i, j+1, dist+1, a, b, min)
		levenshteinDistance(i+1, j+1, dist+1, a, b, min)
	}
}
