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
