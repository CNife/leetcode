package path_with_minimum_effort

import "container/heap"

func MinimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	if n < 2 && m < 2 {
		return 0
	}

	edges := make(edgeHeap, 0, 2*m*n-m-n)

	for i := 0; i < n; i++ {
		for j := 0; j < m-1; j++ {
			left, right := heights[i][j], heights[i][j+1]
			leftId := i*m + j
			rightId := leftId + 1
			heap.Push(&edges, edge{
				from:   leftId,
				to:     rightId,
				weight: diff(left, right),
			})
		}
	}
	for j := 0; j < m; j++ {
		for i := 0; i < n-1; i++ {
			up, down := heights[i][j], heights[i+1][j]
			upId := i*m + j
			downId := upId + m
			heap.Push(&edges, edge{
				from:   upId,
				to:     downId,
				weight: diff(up, down),
			})
		}
	}

	d := newDisjointSet(m * n)
	targetId := m*n - 1
	for len(edges) > 0 {
		edge := heap.Pop(&edges).(edge)
		d.union(edge.from, edge.to)
		if d.find(0) == d.find(targetId) {
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

func diff(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
