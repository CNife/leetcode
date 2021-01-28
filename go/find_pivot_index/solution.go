package find_pivot_index

func PivotIndex(nums []int) int {
	if len(nums) < 3 {
		return -1
	}

	var leftSum, rightSum int
	for _, num := range nums {
		rightSum += num
	}

	for i, num := range nums {
		if i > 0 {
			leftSum += nums[i-1]
		}
		rightSum -= num
		if leftSum == rightSum {
			return i
		}
	}

	return -1
}
