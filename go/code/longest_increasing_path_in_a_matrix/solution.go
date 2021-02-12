package longest_increasing_path_in_a_matrix

func LongestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	if m <= 0 {
		return 0
	}
	n := len(matrix[0])
	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
	}
	ctx := context{matrix, cache, m, n}

	result := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if one := ctx.dfs(i, j); one > result {
				result = one
			}
		}
	}
	return result
}

type context struct {
	matrix [][]int
	cache  [][]int
	m      int
	n      int
}

func (ctx *context) dfs(x, y int) int {
	if cached := ctx.cache[x][y]; cached > 0 {
		return cached
	}

	value := ctx.matrix[x][y]
	result := 0

	if x > 0 && ctx.matrix[x-1][y] > value {
		if up := ctx.dfs(x-1, y); up > result {
			result = up
		}
	}
	if y > 0 && ctx.matrix[x][y-1] > value {
		if right := ctx.dfs(x, y-1); right > result {
			result = right
		}
	}
	if x < ctx.m-1 && ctx.matrix[x+1][y] > value {
		if down := ctx.dfs(x+1, y); down > result {
			result = down
		}
	}
	if y < ctx.n-1 && ctx.matrix[x][y+1] > value {
		if left := ctx.dfs(x, y+1); left > result {
			result = left
		}
	}

	result++
	ctx.cache[x][y] = result
	return result
}
