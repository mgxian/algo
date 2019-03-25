package backtracking

func zeroOneBag(i, cw int, items []int, lw int, mw *int) {
	length := len(items)
	if i == length || cw == lw {
		if cw > *mw {
			*mw = cw
		}
		return
	}

	zeroOneBag(i+1, cw, items, lw, mw)
	if cw+items[i] <= lw {
		zeroOneBag(i+1, cw+items[i], items, lw, mw)
	}
}

func zeroOneBagWithCache(i, cw int, items []int, lw int, mw *int, cache [][]bool) {
	length := len(items)
	if i == length || cw == lw {
		if cw > *mw {
			*mw = cw
		}
		return
	}

	if cache[i][cw] == true {
		return
	}

	cache[i][cw] = true
	zeroOneBagWithCache(i+1, cw, items, lw, mw, cache)
	if cw+items[i] <= lw {
		zeroOneBagWithCache(i+1, cw+items[i], items, lw, mw, cache)
	}
}

func zeroOneBagWithValue(i, cw, cv int, items, value []int, lw int, mv *int) {
	length := len(items)
	if i == length || cw == lw {
		if cv > *mv {
			*mv = cv
		}
		return
	}

	zeroOneBagWithValue(i+1, cw, cv, items, value, lw, mv)
	if cw+items[i] <= lw {
		zeroOneBagWithValue(i+1, cw+items[i], cv+value[i], items, value, lw, mv)
	}
}
