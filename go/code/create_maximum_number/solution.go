package create_maximum_number

func MaxNumber(nums1, nums2 []int, k int) []int {
	var result []int
	for i := 0; i <= k && i <= len(nums1); i++ {
		j := k - i
		if j >= 0 && j <= len(nums2) {
			tmp := merge(subMaxNumber(nums1, i), subMaxNumber(nums2, j))
			if compare(tmp, result, 0, 0) {
				result = tmp
			}
		}
	}
	return result
}

func subMaxNumber(nums []int, length int) []int {
	subNums := make([]int, length)
	current, remain := 0, len(nums)-length
	for i := 0; i < len(nums); i++ {
		for current > 0 && subNums[current-1] < nums[i] && remain > 0 {
			current, remain = current-1, remain-1
		}
		if current < length {
			subNums[current] = nums[i]
			current++
		} else {
			remain--
		}
	}
	return subNums
}

func merge(nums1, nums2 []int) []int {
	length := len(nums1) + len(nums2)
	result := make([]int, length)
	current, p1, p2 := 0, 0, 0
	for current < length {
		if compare(nums1, nums2, p1, p2) {
			result[current] = nums1[p1]
			current, p1 = current+1, p1+1
		} else {
			result[current] = nums2[p2]
			current, p2 = current+1, p2+1
		}
	}
	return result
}

func compare(nums1, nums2 []int, p1, p2 int) bool {
	if p2 >= len(nums2) {
		return true
	}
	if p1 >= len(nums1) {
		return false
	}
	if nums1[p1] > nums2[p2] {
		return true
	}
	if nums1[p1] < nums2[p2] {
		return false
	}
	return compare(nums1, nums2, p1+1, p2+1)
}
