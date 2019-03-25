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
