package backtracking

func lcsLength(i, j int, a, b string) int {
	if i < 0 || j < 0 {
		return 0
	}

	if a[i] == a[j] {
		return 1 + lcsLength(i-1, j-1, a, b)
	}

	n := lcsLength(i-1, j, a, b)
	m := lcsLength(i, j-1, a, b)
	max := n
	if m > max {
		max = m
	}

	return max
}
