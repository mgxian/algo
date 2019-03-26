package dynamic

import (
	"testing"
)

func TestCoinChange(t *testing.T) {
	coins := []int{1, 3, 5}
	amount := 9
	minChangeCount := coinChange(coins, amount)
	t.Log(minChangeCount)

	coins = []int{2, 3, 6}
	amount = 4
	minChangeCount = coinChange(coins, amount)
	t.Log(minChangeCount)
}
