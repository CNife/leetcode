package couples_holding_hands

func MinSwapsCouples(row []int) int {
	n := len(row)
	result := 0
	for i := 0; i < n-1; i += 2 {
		if row[i] == row[i+1]^1 {
			continue
		}
		for j := i + 1; j < n; j++ {
			if row[i] == row[j]^1 {
				row[i+1], row[j] = row[j], row[i+1]
			}
		}
		result++
	}
	return result
}
