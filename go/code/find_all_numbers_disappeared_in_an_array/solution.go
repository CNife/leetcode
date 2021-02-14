package find_all_numbers_disappeared_in_an_array

func FindDisappearedNumbers(nums []int) []int {
	n := len(nums)
	if n < 1 {
		return nil
	}

	for i := range nums {
		index := (nums[i] - 1) % n
		nums[index] += n
	}

	var result []int
	for i := range nums {
		if nums[i] <= n {
			result = append(result, i+1)
		}
	}
	return result
}
