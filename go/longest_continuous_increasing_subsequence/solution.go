package longest_continuous_increasing_subsequence

func FindLengthOfLCIS(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	prev, curLen, maxLen := nums[0], 1, 1
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num > prev && curLen > 0 {
			curLen++
		} else {
			curLen = 1
		}
		if maxLen < curLen {
			maxLen = curLen
		}
		prev = num
	}
	return maxLen
}
