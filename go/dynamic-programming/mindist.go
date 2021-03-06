package dynamic

import "math"

func minDist(w [][]int) (min int) {
	length := len(w)
	states := make([][]int, length)
	for i := range states {
		states[i] = make([]int, length)
	}

	for j := 0; j < length; j++ {
		if j == 0 {
			states[0][j] = w[0][j]
		} else {
			states[0][j] = states[0][j-1] + w[0][j]
		}
	}

	for i := 1; i < length; i++ {
		states[i][0] = states[i-1][0] + w[i][0]
	}

	for i := 1; i < length; i++ {
		for j := 1; j < length; j++ {
			m := states[i][j-1]
			if m > states[i-1][j] {
				m = states[i-1][j]
			}
			states[i][j] = m + w[i][j]
		}
	}

	min = states[length-1][length-1]
	return
}

func minDist2(i, j int, w [][]int, cache [][]int) (min int) {
	if i == 0 && j == 0 {
		return w[0][0]
	}

	if cache[i][j] > 0 {
		return cache[i][j]
	}

	minLeft := math.MaxInt64
	if j-1 >= 0 {
		minLeft = minDist2(i, j-1, w, cache)
	}

	minUp := math.MaxInt64
	if i-1 >= 0 {
		minUp = minDist2(i-1, j, w, cache)
	}

	min = minLeft
	if min > minUp {
		min = minUp
	}

	min += w[i][j]
	cache[i][j] = min
	return
}
