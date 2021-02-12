package insert_interval

func Insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	i, n := 0, len(intervals)
	for ; i < n && intervals[i][1] < newInterval[0]; i++ {
		result = append(result, intervals[i])
	}
	for ; i < n && (intervals[i][1] >= newInterval[0] && intervals[i][0] <= newInterval[1]); i++ {
		newInterval = merge(newInterval, intervals[i])
	}
	result = append(result, newInterval)
	for ; i < n; i++ {
		result = append(result, intervals[i])
	}
	return result
}

func merge(lhs, rhs []int) []int {
	left := min(lhs[0], rhs[0])
	right := max(lhs[1], rhs[1])
	return []int{left, right}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
