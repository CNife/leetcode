package find_first_and_last_position_of_element_in_sorted_array

import "sort"

func SearchRange(nums []int, target int) []int {
	left := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	if left < len(nums) && nums[left] == target {
		right := sort.Search(len(nums), func(i int) bool {
			return nums[i] > target
		})
		return []int{left, right - 1}
	} else {
		return []int{-1, -1}
	}
}
