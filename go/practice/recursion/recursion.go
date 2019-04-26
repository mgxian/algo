package recursion

import (
	"fmt"
)

// Fib returns fib result of number n.
func Fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}

// Factorial returns factorial of number of n
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func swap(data []int, i, j int) {
	data[i], data[j] = data[j], data[i]
}

// Permutation returns number of full permutation
func Permutation(data []int, start, end int) {
	if len(data) == 0 {
		return
	}

	if start == end {
		fmt.Println(data)
		return
	}

	for i := start; i <= end; i++ {
		swap(data, start, i)
		Permutation(data, start+1, end)
		swap(data, start, i)
	}
}
