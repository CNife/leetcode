package patching_array

func MinPatches(nums []int, n int) int {
	total, count, index := 0, 0, 0
	for total < n {
		if index < len(nums) && nums[index] <= total+1 {
			total += nums[index]
			index++
		} else {
			total = 2*total + 1
			count++
		}
	}
	return count
}
