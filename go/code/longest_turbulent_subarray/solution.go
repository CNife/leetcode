package longest_turbulent_subarray

func MaxTurbulenceSize(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	dp := make([]dpItem, n)
	for i := 0; i < n; i++ {
		dp[i] = dpItem{1, 1}
	}

	result := 0
	for i := 1; i < n; i++ {
		if nums[i-1] > nums[i] {
			dp[i].down = dp[i-1].up + 1
		}
		if nums[i-1] < nums[i] {
			dp[i].up = dp[i-1].down + 1
		}
		result = max(result, max(dp[i].up, dp[i].down))
	}
	return result
}

type dpItem struct {
	up, down int
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
