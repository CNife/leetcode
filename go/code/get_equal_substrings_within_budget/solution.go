package get_equal_substrings_within_budget

func EqualSubstring(s string, t string, maxCost int) int {
	left, sumCost := 0, 0
	for right := 0; right < len(s); right++ {
		sumCost += abs(int(s[right]) - int(t[right]))
		if sumCost > maxCost {
			sumCost -= abs(int(s[left]) - int(t[left]))
			left++
		}
	}
	return len(s) - left
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
