package surface_area_of_3d_shapes

func SurfaceArea(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	result := 0
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			height := grid[x][y]
			if height > 0 {
				result += 2
			}
			if x == 0 {
				result += height
			} else if height > grid[x-1][y] {
				result += height - grid[x-1][y]
			}
			if x == m-1 {
				result += height
			} else if height > grid[x+1][y] {
				result += height - grid[x+1][y]
			}
			if y == 0 {
				result += height
			} else if height > grid[x][y-1] {
				result += height - grid[x][y-1]
			}
			if y == n-1 {
				result += height
			} else if height > grid[x][y+1] {
				result += height - grid[x][y+1]
			}
		}
	}
	return result
}
