package intersection_of_two_arrays

import "sort"

func Intersection(arr1, arr2 []int) []int {
	sort.Ints(arr1)
	sort.Ints(arr2)
	var result []int
	for i, j := 0, 0; i < len(arr1) && j < len(arr2); {
		lhs, rhs := arr1[i], arr2[j]
		if lhs < rhs {
			i++
		} else if lhs > rhs {
			j++
		} else {
			if len(result) == 0 || result[len(result)-1] != lhs {
				result = append(result, lhs)
			}
			i++
			j++
		}
	}
	return result
}
