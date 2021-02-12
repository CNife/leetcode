package number_of_operations_to_make_network_connected

func MakeConnected(n int, connections [][]int) int {
	d := newDisjointSet(n)
	subGraphCount, redundantCount := n, 0
	for _, conn := range connections {
		if d.union(conn[0], conn[1]) {
			subGraphCount--
		} else {
			redundantCount++
		}
	}

	if subGraphCount-1 > redundantCount {
		return -1
	} else {
		return subGraphCount - 1
	}
}

type disjointSet []int

func newDisjointSet(size int) disjointSet {
	d := make([]int, size)
	for i := range d {
		d[i] = i
	}
	return d
}

func (d disjointSet) find(x int) int {
	if x != d[x] {
		d[x] = d.find(d[x])
	}
	return d[x]
}

func (d disjointSet) union(x, y int) bool {
	if x != y {
		x, y = d.find(x), d.find(y)
		if x != y {
			d[x] = y
			return true
		}
	}
	return false
}
