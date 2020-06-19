package rotting_oranges

type point struct {
	x int
	y int
}

func OrangesRotting(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rottenOranges := findRottenOranges(grid)
	runs := -1
	for len(rottenOranges) > 0 {
		rottenOranges = rot(grid, rottenOranges)
		runs++
	}

	if hasFreshOranges(grid) {
		return -1
	} else if runs < 0 {
		return 0
	} else {
		return runs
	}
}

func findRottenOranges(grid [][]int) []point {
	var result []point
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				result = append(result, point{i, j})
			}
		}
	}
	return result
}

func rot(grid [][]int, rotten []point) []point {
	m, n := len(grid), len(grid[0])
	var result []point
	for _, p := range rotten {
		if p.x > 0 && grid[p.x-1][p.y] == 1 {
			result = append(result, point{p.x - 1, p.y})
			grid[p.x-1][p.y] = 2
		}
		if p.x < m-1 && grid[p.x+1][p.y] == 1 {
			result = append(result, point{p.x + 1, p.y})
			grid[p.x+1][p.y] = 2
		}
		if p.y > 0 && grid[p.x][p.y-1] == 1 {
			result = append(result, point{p.x, p.y - 1})
			grid[p.x][p.y-1] = 2
		}
		if p.y < n-1 && grid[p.x][p.y+1] == 1 {
			result = append(result, point{p.x, p.y + 1})
			grid[p.x][p.y+1] = 2
		}
	}
	return result
}

func hasFreshOranges(grid [][]int) bool {
	for _, row := range grid {
		for _, cell := range row {
			if cell == 1 {
				return true
			}
		}
	}
	return false
}
