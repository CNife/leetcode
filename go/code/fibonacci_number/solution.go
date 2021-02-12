package fibonacci_number

func Fib(n int) int {
	if n <= 0 {
		return 0
	} else if n < 2 {
		return 1
	}

	a, b := 0, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return b
}
