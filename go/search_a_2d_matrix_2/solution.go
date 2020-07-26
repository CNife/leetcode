package search_a_2d_matrix_2

func SearchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m <= 0 {
		return false
	}
	n := len(matrix[0])
	if n <= 0 {
		return false
	}

	i, j := 0, n-1
	for i >= 0 && i < m && j >= 0 && j < n {
		value := matrix[i][j]
		if value == target {
			return true
		} else if value > target {
			j--
		} else {
			i++
		}
	}
	return false
}
