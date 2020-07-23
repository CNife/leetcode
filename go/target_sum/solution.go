package target_sum

func FindTargetSumWays(nums []int, target int) int {
	result := 0
	backtrack(0, 0, nums, target, &result)
	return result
}

func backtrack(i, sum int, nums []int, target int, result *int) {
	if i == len(nums) {
		if sum == target {
			*result++
		}
	} else {
		backtrack(i+1, sum+nums[i], nums, target, result)
		backtrack(i+1, sum-nums[i], nums, target, result)
	}
}
