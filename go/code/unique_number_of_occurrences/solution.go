package unique_number_of_occurrences

func UniqueOccurrences(array []int) bool {
	m := make(map[int]int)
	for _, num := range array {
		if count, exists := m[num]; exists {
			m[num] = count + 1
		} else {
			m[num] = 1
		}
	}

	counts := make(map[int]struct{})
	for _, count := range m {
		if _, exists := counts[count]; exists {
			return false
		} else {
			counts[count] = struct{}{}
		}
	}
	return true
}
