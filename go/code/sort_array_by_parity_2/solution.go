package sort_array_by_parity_2

func SortArrayByParity(array []int) []int {
	result := make([]int, len(array))
	j, k := 0, 1
	for _, num := range array {
		if num&1 == 0 {
			result[j] = num
			j += 2
		} else {
			result[k] = num
			k += 2
		}
	}
	return result
}
