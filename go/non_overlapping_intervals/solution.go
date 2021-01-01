package non_overlapping_intervals

import "sort"

func EraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) < 1 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	result := 0
	current := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if current[1] <= intervals[i][0] {
			current = intervals[i]
		} else {
			result++
		}
	}
	return result
}
