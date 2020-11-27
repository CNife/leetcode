package four_sum_2

func FourSumCount(a, b, c, d []int) int {
	m := make(map[int]int)
	for _, aNum := range a {
		for _, bNum := range b {
			sum := aNum + bNum
			if c, exists := m[sum]; exists {
				m[sum] = c + 1
			} else {
				m[sum] = 1
			}
		}
	}

	result := 0
	for _, cNum := range c {
		for _, dNum := range d {
			sum := -(cNum + dNum)
			if c, exists := m[sum]; exists {
				result += c
			}
		}
	}
	return result
}
