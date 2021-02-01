package fair_candy_swap

import "sort"

func FairCandySwap(a, b []int) []int {
	sort.Ints(a)
	sort.Ints(b)

	wantDiff := (sum(a) - sum(b)) / 2
	for i, j := 0, 0; i < len(a) && j < len(b); {
		diff := a[i] - b[j]
		if diff == wantDiff {
			return []int{a[i], b[j]}
		} else if diff < wantDiff {
			i++
		} else {
			j++
		}
	}

	panic("unreachable!")
}

func sum(nums []int) int {
	var result int
	for _, num := range nums {
		result += num
	}
	return result
}
