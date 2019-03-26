package dynamic

func double11Adavance(items []int, quota int) (ret []int) {
	length := len(items)
	limitedQuota := 3 * quota
	states := make([][]bool, length)
	for i := range states {
		states[i] = make([]bool, limitedQuota+1)
	}
	states[0][0] = true
	states[0][items[0]] = true

	for i := 1; i < length; i++ {
		for j := 0; j <= limitedQuota; j++ {
			if states[i-1][j] == true {
				states[i][j] = states[i-1][j]
			}
		}

		for j := 0; j <= limitedQuota-items[i]; j++ {
			if states[i-1][j] == true {
				states[i][j+items[i]] = states[i-1][j]
			}
		}
	}

	k := -1
	for j := quota; j <= limitedQuota; j++ {
		if states[length-1][j] == true {
			k = j
			break
		}
	}

	if k == -1 {
		return
	}

	for i := length - 1; i > 0; i-- {
		if k-items[i] >= 0 && states[i-1][k-items[i]] == true {
			ret = append(ret, items[i])
			k -= items[i]
		}
	}

	if k != 0 {
		ret = append(ret, items[0])
	}

	return
}
