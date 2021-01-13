package redundant_connection

func FindRedundantConnection(edges [][]int) []int {
	var d disjointSet
	for _, edge := range edges {
		x, y := find(&d, edge[0]), find(&d, edge[1])
		if x == y {
			return edge
		} else {
			d[x] = y
		}
	}
	return nil
}

type disjointSet []int

func find(d *disjointSet, x int) int {
	for i := len(*d); i <= x; i++ {
		*d = append(*d, i)
	}

	if x != (*d)[x] {
		(*d)[x] = find(d, (*d)[x])
	}
	return (*d)[x]
}
