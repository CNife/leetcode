package minimum_path_sum

func MinPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	dp := initDp(grid[0])
	for i := 1; i < m; i++ {
		row := grid[i]
		dp[0] += row[0]
		for j := 1; j < n; j++ {
			dp[j] = minInt(dp[j], dp[j-1]) + row[j]
		}
	}
	return dp[n-1]
}

func initDp(row []int) []int {
	n := len(row)
	dp := make([]int, n)
	last := 0
	for i := 0; i < n; i++ {
		num := row[i]
		last += num
		dp[i] = last
	}
	return dp
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
