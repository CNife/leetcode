package unique_paths

func UniquePaths(m, n int) int {
	a, b := m-1, n-1
	if a+b <= uint64FactorialBound {
		return fastPath(a, b)
	} else if a < b {
		return slowPath(a, b)
	} else {
		return slowPath(b, a)
	}
}

func fastPath(m, n int) int {
	result := factorials[m+n] / (factorials[m] * factorials[n])
	return int(result)
}

func slowPath(smaller, bigger int) int {
	smallerFactorial := factorials[smaller]
	var acc uint64 = 1
	for i := uint64(bigger + 1); i <= uint64(smaller+bigger); i++ {
		acc *= i
	}
	return int(acc / smallerFactorial)
}

const uint64FactorialBound = 21

var factorials = [uint64FactorialBound]uint64{
	1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800, 39916800, 479001600,
	6227020800, 87178291200, 1307674368000, 20922789888000, 355687428096000,
	6402373705728000, 121645100408832000, 2432902008176640000,
}
