package find_all_continuous_sequences

func findContinuousSequence(target int) [][]int {
	i, j := 1, 1
	sum := 0
	var result [][]int
	for i <= target/2 {
		if sum < target {
			sum += j
			j++
		} else if sum > target {
			sum -= i
			i++
		} else {
			seq := make([]int, j-i)
			for k := i; k < j; k++ {
				seq[k-i] = k
			}
			result = append(result, seq)
			sum -= i
			i++
		}
	}
	return result
}
