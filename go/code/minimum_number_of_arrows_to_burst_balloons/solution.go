package minimum_number_of_arrows_to_burst_balloons

import "sort"

func FindMinArrowShots(points [][]int) int {
	if len(points) < 2 {
		return len(points)
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	interval, result := points[0], 1
	for i := 1; i < len(points); i++ {
		if interval[1] < points[i][0] {
			interval = points[i]
			result++
		} else {
			interval[0] = points[i][0]
			interval[1] = min(interval[1], points[i][1])
		}
	}
	return result
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
