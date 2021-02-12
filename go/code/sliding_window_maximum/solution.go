package sliding_window_maximum

import "math"

func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) < 1 || k < 1 {
		return nil
	}
	if k == 1 {
		return nums
	}
	if k >= len(nums) {
		return []int{maxOf(nums)}
	}

	left := leftMax(nums, k)
	right := rightMax(nums, k)
	return buildMax(left, right, k)
}

func maxOf(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = max(result, nums[i])
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func leftMax(nums []int, k int) []int {
	result := make([]int, len(nums))
	i, blockMax := 0, math.MinInt32
	for j, num := range nums {
		if i >= k {
			i = 0
			blockMax = num
		}
		blockMax = max(blockMax, num)
		result[j] = blockMax
		i++
	}
	return result
}

func rightMax(nums []int, k int) []int {
	result := make([]int, len(nums))
	i, blockMax := len(nums)%k, math.MinInt32
	for j := len(nums) - 1; j >= 0; j-- {
		if i == 0 {
			i = k
			blockMax = nums[j]
		}
		blockMax = max(blockMax, nums[j])
		result[j] = blockMax
		i--
	}
	return result
}

func buildMax(left []int, right []int, k int) []int {
	result := make([]int, len(left)-k+1)
	for i := 0; i <= len(left)-k; i++ {
		j := i + k - 1
		result[i] = max(left[j], right[i])
	}
	return result
}
