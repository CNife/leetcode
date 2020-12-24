package rotate_image

func Rotate(matrix [][]int) {
	n := len(matrix) - 1
	for r := 0; r < (n+1)/2; r++ {
		for c := r; c < n-r; c++ {
			matrix[r][c], matrix[n-c][r], matrix[n-r][n-c], matrix[c][n-r] =
				matrix[n-c][r], matrix[n-r][n-c], matrix[c][n-r], matrix[r][c]
		}
	}
}
