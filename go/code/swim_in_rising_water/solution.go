package swim_in_rising_water

import "container/heap"

func SwimInWater(grid [][]int) int {
	n := len(grid)
	if n < 2 {
		return 0
	}

	eh := make(edgeHeap, 0, 2*n*(n-1))
	for i, row := range grid {
		for j := 0; j < n-1; j++ {
			left, right := row[j], row[j+1]
			leftId := i*n + j
			heap.Push(&eh, edge{leftId, leftId + 1, max(left, right)})
		}
	}
	for j := 0; j < n; j++ {
		for i := 0; i < n-1; i++ {
			up, down := grid[i][j], grid[i+1][j]
			upId := i*n + j
			heap.Push(&eh, edge{upId, upId + n, max(up, down)})
		}
	}

	d := newDisjointSet(n * n)
	for len(eh) > 0 {
		edge := heap.Pop(&eh).(edge)
		d.union(edge.from, edge.to)
		if d.find(0) == d.find(n*n-1) {
			return edge.weight
		}
	}
	panic("unreachable")
}

type edge struct {
	from, to, weight int
}

type edgeHeap []edge

func (eh edgeHeap) Len() int {
	return len(eh)
}

func (eh edgeHeap) Less(i, j int) bool {
	return eh[i].weight < eh[j].weight
}

func (eh edgeHeap) Swap(i, j int) {
	eh[i], eh[j] = eh[j], eh[i]
}

func (eh *edgeHeap) Push(x interface{}) {
	*eh = append(*eh, x.(edge))
}

func (eh *edgeHeap) Pop() interface{} {
	old := *eh
	n := len(old)
	x := old[n-1]
	*eh = old[:n-1]
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
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

func (d disjointSet) union(i, j int) {
	if i != j {
		i, j = d.find(i), d.find(j)
		if i != j {
			d[i] = j
		}
	}
}
