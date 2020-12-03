package count_primes

import "math"

func CountPrimes(n int) int {
	if n < 3 {
		return 0
	}

	nums := make([]bool, n)
	for i := range nums {
		nums[i] = true
	}
	m := int(math.Floor(math.Sqrt(float64(n))))
	for i := 2; i <= m; i++ {
		if nums[i] {
			for j := i * i; j < n; j = j + i {
				nums[j] = false
			}
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if nums[i] {
			count++
		}
	}
	return count
}
