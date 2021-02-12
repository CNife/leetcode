package reverse_pairs

func ReversePairs(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	sorted := make([]int, len(nums))
	return mergeSort(nums, sorted, 0, len(nums)-1)
}

func mergeSort(nums, sorted []int, left, right int) int {
	if left == right {
		return 0
	}

	middle := left + (right-left)/2
	result := mergeSort(nums, sorted, left, middle) +
		mergeSort(nums, sorted, middle+1, right) +
		reversedPairs(nums, left, right)

	i, j, k := left, middle+1, left
	for i <= middle && j <= right {
		if nums[i] <= nums[j] {
			sorted[k] = nums[i]
			k, i = k+1, i+1
		} else {
			sorted[k] = nums[j]
			k, j = k+1, j+1
		}
	}
	for i <= middle {
		sorted[k] = nums[i]
		k, i = k+1, i+1
	}
	for j <= right {
		sorted[k] = nums[j]
		k, j = k+1, j+1
	}

	for idx := left; idx <= right; idx++ {
		nums[idx] = sorted[idx]
	}
	return result
}

func reversedPairs(nums []int, left, right int) int {
	result, middle := 0, left+(right-left)/2
	for i, j := left, middle+1; i <= middle; i++ {
		for j <= right && nums[i] > 2*nums[j] {
			result += middle - i + 1
			j++
		}
	}
	return result
}
