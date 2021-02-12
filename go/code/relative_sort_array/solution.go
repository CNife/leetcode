package relative_sort_array

import "sort"

func RelativeSortArray(arr1, arr2 []int) []int {
	baseKey := len(arr2)
	cmpMap := make(map[int]int)
	for i, elem := range arr2 {
		cmpMap[elem] = i
	}
	cmpKey := func(num int) int {
		if idx, exists := cmpMap[num]; exists {
			return idx
		} else {
			return baseKey + num
		}
	}
	sort.Slice(arr1, func(i, j int) bool {
		return cmpKey(arr1[i]) < cmpKey(arr1[j])
	})
	return arr1
}
