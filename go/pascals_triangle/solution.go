package pascals_triangle

func Generate(numRows int) [][]int {
	if numRows < 1 {
		return nil
	}

	result := make([][]int, numRows)
	result[0] = []int{1}
	for i := 1; i < numRows; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = result[i-1][j] + result[i-1][j-1]
		}
		result[i] = row
	}
	return result
}
