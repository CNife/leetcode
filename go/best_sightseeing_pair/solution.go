package best_sightseeing_pair

func MaxScoreSightseeingPair(array []int) int {
	result, mx := 0, array[0]
	for i := 1; i < len(array); i++ {
		result = max(result, mx+array[i]-i)
		mx = max(mx, array[i]+i)
	}
	return result
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
