package pascals_triangle_2

func GetRow(rowIndex int) []int {
	if rowIndex < 0 {
		panic("rowIndex < 0")
	}
	if rowIndex == 0 {
		return []int{1}
	}

	result := make([]int, rowIndex+1)
	result[0], result[1] = 1, 1

	for i := 1; i < rowIndex; i++ {
		result[i+1] = 1
		for j := i; j >= 1; j-- {
			result[j] += result[j-1]
		}
	}
	return result
}
