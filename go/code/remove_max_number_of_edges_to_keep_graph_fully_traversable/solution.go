package remove_max_number_of_edges_to_keep_graph_fully_traversable

func MaxNumEdgesToRemove(n int, edges [][]int) int {
	result := 0

	d3 := newDisjointSet(n + 1)
	for _, edge := range edges {
		if edge[0] == 3 && !d3.union(edge[1], edge[2]) {
			result++
		}
	}

	d1, d2 := d3, cloneDisjointSet(d3)
	for _, edge := range edges {
		switch edge[0] {
		case 1:
			if !d1.union(edge[1], edge[2]) {
				result++
			}
		case 2:
			if !d2.union(edge[1], edge[2]) {
				result++
			}
		}
	}

	if d1.counter != 1 || d2.counter != 1 {
		return -1
	}
	return result
}

type disjointSet struct {
	inner   []int
	counter int
}

func newDisjointSet(size int) disjointSet {
	inner := make([]int, size)
	for i := range inner {
		inner[i] = i
	}
	return disjointSet{inner, size - 1}
}

func cloneDisjointSet(src disjointSet) disjointSet {
	inner := make([]int, len(src.inner))
	copy(inner, src.inner)
	return disjointSet{inner, src.counter}
}

func (d *disjointSet) find(i int) int {
	p := &d.inner[i]
	if *p != i {
		*p = d.find(*p)
	}
	return *p
}

func (d *disjointSet) union(i, j int) bool {
	if i != j {
		i, j = d.find(i), d.find(j)
		if i != j {
			d.inner[i] = j
			d.counter--
			return true
		}
	}
	return false
}
