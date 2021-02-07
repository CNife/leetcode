package non_decreasing_array

func CheckPossibility(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	flag := nums[0] <= nums[1]
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if flag {
				if nums[i+1] >= nums[i-1] {
					nums[i] = nums[i+1]
				} else {
					nums[i+1] = nums[i]
				}
				flag = false
			} else {
				return false
			}
		}
	}
	return true
}
