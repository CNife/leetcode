package triangle

import "math"

func MinimumTotal(triangle [][]int) int {
	dp := make([]int, len(triangle))
	for _, row := range triangle {
		for i := len(row) - 1; i >= 0; i-- {
			var prevMin int
			switch i {
			case 0:
				prevMin = dp[0]
			case len(row) - 1:
				prevMin = dp[i-1]
			default:
				prevMin = min(dp[i-1], dp[i])
			}
			dp[i] = prevMin + row[i]
		}
	}

	result := math.MaxInt32
	for _, dpValue := range dp {
		result = min(result, dpValue)
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
