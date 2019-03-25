package dynamic

func zeroOneBag(items []int, limitedWeight int) int {
	length := len(items)
	states := make([][]bool, length)
	for i := range states {
		states[i] = make([]bool, limitedWeight+1)
	}

	states[0][0] = true
	states[0][items[0]] = true
	for i := 1; i < length; i++ {
		for j := 0; j <= limitedWeight; j++ {
			if states[i-1][j] == true {
				states[i][j] = true
			}
		}

		for j := 0; j <= limitedWeight-items[i]; j++ {
			if states[i-1][j] == true {
				states[i][j+items[i]] = true
			}
		}
	}

	for i := limitedWeight; i >= 0; i-- {
		if states[length-1][i] == true {
			return i
		}
	}

	return -1
}

func zeroOneBag2(items []int, limitedWeight int) int {
	length := len(items)
	states := make([]bool, limitedWeight+1)
	states[0] = true
	states[items[0]] = true
	for i := 1; i < length; i++ {
		for j := limitedWeight - items[i]; j >= 0; j-- {
			if states[j] == true {
				states[j+items[i]] = true
			}
		}
	}

	for i := limitedWeight; i >= 0; i-- {
		if states[i] == true {
			return i
		}
	}

	return -1
}

func zeroOneBagWithValue(items, value []int, limitedWeight int) int {
	length := len(items)
	states := make([][]int, length)
	for i := range states {
		states[i] = make([]int, limitedWeight+1)
		for j := range states[i] {
			states[i][j] = -1
		}
	}

	states[0][0] = 0
	states[0][items[0]] = value[0]
	for i := 1; i < length; i++ {
		for j := 0; j <= limitedWeight; j++ {
			if states[i-1][j] > -1 {
				states[i][j] = states[i-1][j]
			}
		}

		for j := 0; j <= limitedWeight-items[i]; j++ {
			if states[i-1][j] > -1 {
				v := states[i-1][j] + value[i]
				if v > states[i][j+items[i]] {
					states[i][j+items[i]] = v
				}
			}
		}
	}

	maxValue := -1
	for j := 0; j <= limitedWeight; j++ {
		if states[length-1][j] > maxValue {
			maxValue = states[length-1][j]
		}
	}
	return maxValue
}
