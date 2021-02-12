package smallest_string_with_swaps

import "sort"

func SmallestStringWithSwaps(s string, pairs [][]int) string {
	if len(s) < 2 {
		return s
	}

	d := newDisjointSet(len(s))
	for _, pair := range pairs {
		d.union(pair[0], pair[1])
	}

	cm := map[int][]int{}
	for i := 0; i < len(s); i++ {
		key := d.find(i)
		if indices, exists := cm[key]; exists {
			cm[key] = append(indices, i)
		} else {
			cm[key] = []int{i}
		}
	}

	buffer := make([]byte, len(s))
	for _, indices := range cm {
		bytes := make([]byte, len(indices))
		for i, index := range indices {
			bytes[i] = s[index]
		}
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})

		for i, index := range indices {
			buffer[index] = bytes[i]
		}
	}
	return string(buffer)
}

type disjointSet []int

func newDisjointSet(size int) disjointSet {
	items := make([]int, size)
	for i := range items {
		items[i] = i
	}
	return items
}

func (d disjointSet) find(x int) int {
	if x != d[x] {
		d[x] = d.find(d[x])
	}
	return d[x]
}

func (d disjointSet) union(x, y int) {
	x, y = d.find(x), d.find(y)
	d[x] = y
}
