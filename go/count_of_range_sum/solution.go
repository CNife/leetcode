package count_of_range_sum

import "sort"

func CountRangeSum(nums []int, lower, higher int) int {
	result, now, pre := 0, 0, []int{0}
	for _, n := range nums {
		now += n

		left := sort.Search(len(pre), func(i int) bool { return pre[i] >= now-higher })
		right := sort.Search(len(pre), func(i int) bool { return pre[i] > now-lower })
		result += right - left

		insertIdx := sort.Search(len(pre), func(i int) bool { return pre[i] > now })
		pre = insert(pre, insertIdx, now)
	}
	return result
}

func insert(nums []int, idx, x int) []int {
	n := len(nums)
	if idx < 0 || idx > n {
		panic("index out of bound")
	}
	if idx == n {
		return append(nums, x)
	} else {
		nums = append(nums[:idx+1], nums[idx:]...)
		nums[idx] = x
		return nums
	}
}
