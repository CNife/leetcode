package ugly_number_2

func NthUglyNumber(n int) int {
	if n < 1 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = 1

	p2, p3, p5 := 0, 0, 0
	for i := 1; i < n; i++ {
		n2, n3, n5 := 2*dp[p2], 3*dp[p3], 5*dp[p5]
		dp[i] = min3(n2, n3, n5)
		if dp[i] == n2 {
			p2++
		}
		if dp[i] == n3 {
			p3++
		}
		if dp[i] == n5 {
			p5++
		}
	}
	return dp[n-1]
}

func min3(x, y, z int) int {
	if x < y {
		if x < z {
			return x
		}
		return z
	}
	if y < z {
		return y
	}
	return z
}
