package stitch_videos

func VideoStitching(clips [][]int, T int) int {
	dp := make([]int, T+1)
	dp[0] = 0
	for i := 1; i <= T; i++ {
		dp[i] = 1000
	}

	for i := 1; i <= T; i++ {
		for _, clip := range clips {
			start, end := clip[0], clip[1]
			if start < i && i <= end {
				dp[i] = min(dp[i], dp[start]+1)
			}
		}
	}

	if dp[T] == 1000 {
		return -1
	}
	return dp[T]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
