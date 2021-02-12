package score_after_flipping_matrix

func MatrixScore(matrix [][]int) int {
	if len(matrix) <= 0 || len(matrix[0]) <= 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	result := m * (1 << (n - 1))
	for i := 1; i < n; i++ {
		count := 0
		for j := 0; j < m; j++ {
			if matrix[j][i] == matrix[j][0] {
				count++
			}
		}
		result += max(count, m-count) * (1 << (n - 1 - i))
	}
	return result
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
