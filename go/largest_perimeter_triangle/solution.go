package largest_perimeter_triangle

import "sort"

func LargestPerimeter(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[j] < nums[i]
	})
	first, second, third := 0, 1, 2
	for third < len(nums) {
		if nums[third]+nums[second] > nums[first] {
			return nums[first] + nums[second] + nums[third]
		} else {
			first, second, third = first+1, second+1, third+1
		}
	}
	return 0
}
