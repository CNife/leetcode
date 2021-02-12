package maximum_product_of_three_numbers

import "sort"

func MaximumProduct(nums []int) int {
	sort.Ints(nums)

	n := len(nums)
	max3Product := nums[n-3] * nums[n-2] * nums[n-1]
	if nums[0] < 0 && nums[1] < 0 {
		if another := nums[0] * nums[1] * nums[n-1]; another > max3Product {
			return another
		}
	}
	return max3Product
}
