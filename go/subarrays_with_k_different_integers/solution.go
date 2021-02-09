package subarrays_with_k_different_integers

func SubarraysWithKDistinct(nums []int, k int) int {
	return getMostDistinct(nums, k) - getMostDistinct(nums, k-1)
}

func getMostDistinct(nums []int, k int) int {
	var left, right, result int
	counter := map[int]int{}

	for right < len(nums) {
		rightNum := nums[right]
		if count, exists := counter[rightNum]; exists {
			counter[rightNum] = count + 1
		} else {
			counter[rightNum] = 1
		}
		right++

		for k < len(counter) {
			leftNum := nums[left]
			if count := counter[leftNum]; count > 1 {
				counter[leftNum] = count - 1
			} else {
				delete(counter, leftNum)
			}
			left++
		}

		result += right - left
	}

	return result
}
