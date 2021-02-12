package check_if_it_is_a_straight_line

import "math"

var epsilon = math.Nextafter(1.0, 2.0) - 1.0

func CheckStraightLine(coordinates [][]int) bool {
	if len(coordinates) < 2 {
		return true
	}

	first, second := coordinates[0], coordinates[1]
	if first[0] == second[0] {
		for i := 2; i < len(coordinates); i++ {
			if coordinates[i][0] != first[0] {
				return false
			}
		}
		return true
	}

	k := float64(first[1]-second[1]) / float64(first[0]-second[0])
	d := float64(first[1]) - k*float64(first[0])
	for i := 2; i < len(coordinates); i++ {
		x, y := float64(coordinates[i][0]), float64(coordinates[i][1])
		wantY := x*k + d
		if wantY > y {
			wantY, y = y, wantY
		}
		if y-wantY >= epsilon {
			return false
		}
	}
	return true
}
