package matrix_cells_in_distance_order

func AllCellsDistanceOrder(r, c, r0, c0 int) [][]int {
	var distPoints [][][]int
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			dist := distance(i, j, r0, c0)
			if dist >= len(distPoints) {
				for k := len(distPoints); k <= dist; k++ {
					distPoints = append(distPoints, nil)
				}
			}
			//goland:noinspection GoNilness
			distPoints[dist] = append(distPoints[dist], []int{i, j})
		}
	}

	var result [][]int
	for _, ps := range distPoints {
		result = append(result, ps...)
	}
	return result
}

func distance(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
