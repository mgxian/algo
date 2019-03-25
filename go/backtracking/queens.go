package backtracking

import "fmt"

func nQueens(row int, ret []int) {
	if row == len(ret) {
		printQueen(ret)
		return
	}

	for column := 0; column < len(ret); column++ {
		if isOK(row, column, ret) {
			ret[row] = column
			nQueens(row+1, ret)
		}
	}
}

func isOK(row, column int, ret []int) bool {
	leftUp := column - 1
	rightUp := column + 1
	for k := row - 1; k >= 0; k-- {
		if ret[k] == column {
			return false
		}

		if leftUp >= 0 && ret[k] == leftUp {
			return false
		}

		if rightUp < len(ret) && ret[k] == rightUp {
			return false
		}

		leftUp--
		rightUp++
	}
	return true
}

func printQueen(ret []int) {
	result := "One Way to Put Queue:\n"
	columns := len(ret)
	for _, column := range ret {
		for i := 0; i < columns; i++ {
			if i == column {
				result += "Q"
			} else {
				result += "."
			}
		}
		result += "\n"
	}
	result += "\n\n"
	fmt.Print(result)
}
