package intersection_of_two_arrays_2

func intersect(lhs []int, rhs []int) []int {
	lhsMap, rhsMap := buildMap(lhs), buildMap(rhs)
	if len(lhsMap) > len(rhsMap) {
		lhsMap, rhsMap = rhsMap, lhsMap
	}

	var result []int
	for num, lhsCount := range lhsMap {
		rhsCount := rhsMap[num]
		resultCount := lhsCount
		if rhsCount < lhsCount {
			resultCount = rhsCount
		}
		for i := 0; i < resultCount; i++ {
			result = append(result, num)
		}
	}
	return result
}

func buildMap(nums []int) map[int]int {
	result := make(map[int]int)
	for _, num := range nums {
		if count, ok := result[num]; ok {
			result[num] = count + 1
		} else {
			result[num] = 1
		}
	}
	return result
}
