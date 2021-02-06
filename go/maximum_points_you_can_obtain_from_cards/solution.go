package maximum_points_you_can_obtain_from_cards

func MaxScore(cardPoints []int, k int) int {
	windowSize := len(cardPoints) - k
	windowSum := 0
	for _, num := range cardPoints[:windowSize] {
		windowSum += num
	}

	minWindowSum, sum := windowSum, windowSum
	for i, num := range cardPoints[windowSize:] {
		windowSum = windowSum + num - cardPoints[i]
		sum += num
		if windowSum < minWindowSum {
			minWindowSum = windowSum
		}
	}
	return sum - minWindowSum
}
