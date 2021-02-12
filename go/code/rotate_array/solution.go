package rotate_array

func Rotate(nums []int, k int) {
	length := len(nums)
	k %= length
	if k > 0 {
		buffer := make([]int, k)
		copy(buffer, nums[length-k:])
		copy(nums[k:], nums[:length-k])
		copy(nums, buffer)
	}
}
