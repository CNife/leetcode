package maximal_square

func MaximalSquare(matrix [][]byte) int {
	if len(matrix) <= 0 || len(matrix[0]) <= 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	if m == 1 {
		for i := 0; i < n; i++ {
			if matrix[0][i] == '1' {
				return 1
			}
		}
		return 0
	}
	if n == 1 {
		for i := 0; i < m; i++ {
			if matrix[i][0] == '1' {
				return 1
			}
		}
		return 0
	}

	dp, result := make([]int, n), 0
	for i := 0; i < n; i++ {
		if matrix[0][i] == '1' {
			result = 1
			dp[i] = 1
		}
	}

	for i := 1; i < m; i++ {
		prevLeft, left := dp[0], 0
		if matrix[i][0] == '1' {
			if result == 0 {
				result = 1
			}
			left = 1
			dp[0] = 1
		} else {
			dp[0] = 0
		}
		for j := 1; j < n; j++ {
			up := dp[j]
			if matrix[i][j] == '1' {
				dp[j] = min3(prevLeft, left, up) + 1
			} else {
				dp[j] = 0
			}
			if dp[j] > result {
				result = dp[j]
			}
			prevLeft, left = up, dp[j]
		}
	}
	return result * result
}

func min3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}
