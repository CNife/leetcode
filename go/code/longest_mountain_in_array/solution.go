package longest_mountain_in_array

func LongestMountain(array []int) int {
	n, maxLen := len(array), 0
	for i := 1; i < n; {
		inc, dec := 0, 0
		for i < n && array[i-1] < array[i] {
			inc++
			i++
		}
		for i < n && array[i-1] > array[i] {
			dec++
			i++
		}
		if inc > 0 && dec > 0 {
			if mountainLen := inc + dec + 1; maxLen < mountainLen {
				maxLen = mountainLen
			}
		}
		for i < n && array[i-1] == array[i] {
			i++
		}
	}
	return maxLen
}
