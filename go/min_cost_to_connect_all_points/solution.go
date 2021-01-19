package min_cost_to_connect_all_points

import "container/heap"

func MinCostConnectPoints(points [][]int) int {
	n := len(points)
	if n < 2 {
		return 0
	}

	var edges edgeHeap
	for i := 0; i < n-1; i++ {
		start := points[i]
		for j := i + 1; j < n; j++ {
			end := points[j]
			heap.Push(&edges, edge{i, j, distance(start, end)})
		}
	}

	d, graphCount := newDisjointSet(n), n
	result := 0
	for graphCount > 1 {
		e := heap.Pop(&edges).(edge)
		startGraph, endGraph := d.find(e.start), d.find(e.end)
		if startGraph != endGraph {
			graphCount--
			d[startGraph] = endGraph
			result += e.distance
		}
	}

	return result
}

func distance(x, y []int) int {
	xDist, yDist := x[0]-y[0], x[1]-y[1]
	if xDist < 0 {
		xDist = -xDist
	}
	if yDist < 0 {
		yDist = -yDist
	}
	return xDist + yDist
}

type edge struct {
	start, end int
	distance   int
}

type edgeHeap []edge

func (e edgeHeap) Len() int {
	return len(e)
}

func (e edgeHeap) Less(i, j int) bool {
	return e[i].distance < e[j].distance
}

func (e edgeHeap) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e *edgeHeap) Push(x interface{}) {
	*e = append(*e, x.(edge))
}

func (e *edgeHeap) Pop() interface{} {
	old := *e
	n := len(old)
	x := old[n-1]
	*e = old[:n-1]
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
