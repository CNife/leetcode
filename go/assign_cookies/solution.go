package assign_cookies

import "sort"

func FindContentChildren(g, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	result := 0
	for i, j := 0, 0; i < len(s) && j < len(g); i++ {
		if g[j] <= s[i] {
			j++
			result++
		}
	}
	return result
}
