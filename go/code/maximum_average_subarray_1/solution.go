package maximum_average_subarray_1

func FindMaxAverage(nums []int, k int) float64 {
	if k == 1 {
		return float64(maxOf(nums))
	}
	if k == len(nums) {
		return float64(sumOf(nums)) / float64(len(nums))
	}

	sliceSum := sumOf(nums[:k])
	maxSum := sliceSum
	for i := k; i < len(nums); i++ {
		sliceSum = sliceSum + nums[i] - nums[i-k]
		if sliceSum > maxSum {
			maxSum = sliceSum
		}
	}
	return float64(maxSum) / float64(k)
}

func maxOf(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

func sumOf(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}
