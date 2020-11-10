package next_permutation

func NextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}

	rFind := func(start, end int, predicate func(int) bool) int {
		for i := end - 1; i >= start; i-- {
			if predicate(i) {
				return i
			}
		}
		return -1
	}

	m := rFind(0, len(nums)-1, func(i int) bool {
		return nums[i] < nums[i+1]
	})
	if m < 0 {
		reverse(nums)
		return
	}

	n := rFind(m+1, len(nums), func(i int) bool {
		return nums[i] > nums[m]
	})
	nums[n], nums[m] = nums[m], nums[n]
	reverse(nums[m+1:])
}

func reverse(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
