package dynamic

func coinChange(coins []int, amount int) int {
	length := len(coins)
	states := make([][]int, length)
	for i := range states {
		states[i] = make([]int, amount+1)
	}

	states[0][0] = 0
	for j := 1; j <= amount; j++ {
		maxChangeCount := j / coins[0]
		remainChange := j - maxChangeCount*coins[0]
		if remainChange == 0 {
			states[0][j] = maxChangeCount
		}
	}

	for i := 1; i < length; i++ {
		for j := 0; j <= amount; j++ {
			maxChangeCount := j / coins[i]
			remainChange := j - maxChangeCount*coins[i]
			if remainChange != 0 && states[i-1][remainChange] == 0 {
				return -1
			}

			v := maxChangeCount + states[i-1][remainChange]
			states[i][j] = v
		}
	}
	return states[length-1][amount]
}
