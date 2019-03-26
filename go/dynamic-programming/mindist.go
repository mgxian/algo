package dynamic

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
