package recursion

// Fib returns fib result of number n.
func Fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}

// Factorial return factorial of number of n
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
