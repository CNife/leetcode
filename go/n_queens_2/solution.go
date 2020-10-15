package n_queens_2

func TotalNQueens(n int) int {
	if n < 1 {
		return 0
	}

	result := 0
	columns := make([]bool, n)
	diagonals, antiDiagonals := make([]bool, n*2-1), make([]bool, n*2-1)

	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			result++
			return
		}
		for j := 0; j < n; j++ {
			diagonalIndex, antiDiagonalIndex := n-1-i+j, i+j
			if columns[j] || diagonals[diagonalIndex] || antiDiagonals[antiDiagonalIndex] {
				continue
			}
			columns[j], diagonals[diagonalIndex], antiDiagonals[antiDiagonalIndex] = true, true, true
			dfs(i + 1)
			columns[j], diagonals[diagonalIndex], antiDiagonals[antiDiagonalIndex] = false, false, false
		}
	}

	dfs(0)
	return result
}
