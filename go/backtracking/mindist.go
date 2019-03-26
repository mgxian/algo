package backtracking

func minDist(i, j, dist int, w [][]int, min *int) {
	length := len(w)
	if i == length-1 && j == length-1 {
		dist += w[length-1][length-1]
		if *min > dist {
			*min = dist
		}
		return
	}

	if i == length || j == length {
		return
	}

	if i < length {
		minDist(i+1, j, dist+w[i][j], w, min)
	}

	if j < length {
		minDist(i, j+1, dist+w[i][j], w, min)
	}
}
