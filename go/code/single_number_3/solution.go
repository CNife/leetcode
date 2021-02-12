package single_number_3

func SingleNumber(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	mask := xor & (-xor)
	result := make([]int, 2)
	for _, num := range nums {
		if num&mask == 0 {
			result[0] ^= num
		} else {
			result[1] ^= num
		}
	}
	return result
}
