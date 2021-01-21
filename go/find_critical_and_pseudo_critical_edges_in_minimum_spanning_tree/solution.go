package find_critical_and_pseudo_critical_edges_in_minimum_spanning_tree

import "container/heap"

func FindCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	mb := newMSTBuilder(n)
	mb.addEdges(cloneEdges(edges))
	mstWeightSum := mb.weightSum

	var criticalEdges, nonCriticalEdges []int
	for i := range edges {
		clonedEdges := cloneEdgesWithout(edges, i)

		mb.reset()
		mb.addEdges(clonedEdges)
		if mb.isMST() && mb.weightSum == mstWeightSum {
			nonCriticalEdges = append(nonCriticalEdges, i)
		} else {
			criticalEdges = append(criticalEdges, i)
		}
	}

	var pseudoEdges []int
	for _, i := range nonCriticalEdges {
		mb.reset()
		mb.addEdge(edges[i])
		mb.addEdges(cloneEdges(edges))
		if mb.isMST() && mb.weightSum == mstWeightSum {
			pseudoEdges = append(pseudoEdges, i)
		}
	}

	return [][]int{criticalEdges, pseudoEdges}
}

type mstBuilder struct {
	d         disjointSet
	counter   int
	edges     [][]int
	weightSum int
}

func newMSTBuilder(n int) *mstBuilder {
	return &mstBuilder{
		d:       newDisjointSet(n),
		counter: n,
	}
}

func (mb *mstBuilder) isMST() bool {
	return mb.counter == 1
}

func (mb *mstBuilder) reset() {
	for i := range mb.d {
		mb.d[i] = i
	}
	mb.counter = len(mb.d)
	mb.edges = mb.edges[:0]
	mb.weightSum = 0
}

func (mb *mstBuilder) addEdge(e []int) {
	from, to, weight := e[0], e[1], e[2]
	from, to = mb.d.find(from), mb.d.find(to)
	if from != to {
		mb.d[from] = to
		mb.counter--
		mb.edges = append(mb.edges, e)
		mb.weightSum += weight
	}
}

func (mb *mstBuilder) addEdges(edges [][]int) {
	h := edgeHeap(edges)
	heap.Init(&h)

	for h.Len() > 0 {
		e := heap.Pop(&h).([]int)
		mb.addEdge(e)
	}
}

type edgeHeap [][]int

func (h edgeHeap) Len() int {
	return len(h)
}

func (h edgeHeap) Less(i, j int) bool {
	return h[i][2] < h[j][2]
}

func (h edgeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *edgeHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *edgeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type disjointSet []int

func newDisjointSet(size int) disjointSet {
	d := make([]int, size)
	for i := 0; i < size; i++ {
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

func cloneEdges(edges [][]int) [][]int {
	cloned := make([][]int, len(edges))
	copy(cloned, edges)
	return cloned
}

func cloneEdgesWithout(edges [][]int, i int) [][]int {
	cloned := make([][]int, len(edges)-1)
	copy(cloned[:i], edges[:i])
	copy(cloned[i:], edges[i+1:])
	return cloned
}
