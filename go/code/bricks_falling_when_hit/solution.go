package bricks_falling_when_hit

func HitBricks(grid [][]int, hits [][]int) []int {
	m, n := len(grid), len(grid[0])

	result := make([]int, len(hits))
	for i := range result {
		result[i] = -1
	}

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
			grid[x][y] = 2
			return 1 + dfs(x+1, y) + dfs(x-1, y) + dfs(x, y-1) + dfs(x, y+1)
		}
		return 0
	}
	isStable := func(x, y int) bool {
		return x == 0 ||
			x < m-1 && grid[x+1][y] == 2 ||
			x > 0 && grid[x-1][y] == 2 ||
			y < n-1 && grid[x][y+1] == 2 ||
			y > 0 && grid[x][y-1] == 2
	}

	for _, hit := range hits {
		x, y := hit[0], hit[1]
		grid[x][y]--
	}

	for y := 0; y < n; y++ {
		dfs(0, y)
	}

	for i := len(hits) - 1; i >= 0; i-- {
		x, y := hits[i][0], hits[i][1]
		grid[x][y]++
		if !isStable(x, y) || grid[x][y] == 0 {
			result[i] = 0
		} else {
			result[i] = dfs(x, y) - 1
		}
	}
	return result
}
