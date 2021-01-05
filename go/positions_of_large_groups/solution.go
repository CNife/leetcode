package positions_of_large_groups

func LargeGroupPosition(s string) [][]int {
	var char byte
	var count, start int
	var result [][]int

	for i := 0; i < len(s); i++ {
		if char == s[i] {
			count++
		} else {
			if count > 2 {
				result = append(result, []int{start, i - 1})
			}
			char, count, start = s[i], 1, i
		}
	}
	if count > 2 {
		result = append(result, []int{start, len(s) - 1})
	}

	return result
}
