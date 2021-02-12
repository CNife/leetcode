package similar_string_groups

import "fmt"

func NumSimilarGroups(strs []string) int {
	d := newDisjointSet(len(strs))
	counter := len(strs)

	for i := 0; i < len(strs)-1; i++ {
		for j := i + 1; j < len(strs); j++ {
			ip, jp := d.find(i), d.find(j)
			if ip != jp && isSimilar(strs[i], strs[j]) {
				d[ip] = jp
				counter--
			}
		}
	}

	fmt.Println(d)
	return counter
}

func isSimilar(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	diffs, k := [2]int{-1, -1}, 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if k < 2 {
				diffs[k] = i
				k++
			} else {
				return false
			}
		}
	}

	return k == 0 || k == 2 &&
		s1[diffs[0]] == s2[diffs[1]] &&
		s1[diffs[1]] == s2[diffs[0]]
}

type disjointSet []int

func newDisjointSet(size int) disjointSet {
	d := make([]int, size)
	for i := range d {
		d[i] = i
	}
	return d
}

func (d disjointSet) find(i int) int {
	if d[i] != i {
		d[i] = d.find(d[i])
	}
	return d[i]
}
