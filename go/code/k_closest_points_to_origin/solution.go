package k_closest_points_to_origin

import "container/heap"

func KClosest(points [][]int, k int) [][]int {
	if k <= 0 {
		return nil
	}
	if len(points) <= k {
		return points
	}

	ph := &pointHeap{points}
	heap.Init(ph)

	var result [][]int
	for i := 0; i < k; i++ {
		p := heap.Pop(ph)
		result = append(result, p.([]int))
	}
	return result
}

type pointHeap struct {
	points [][]int
	//distances []uint64
}

func (ph *pointHeap) Len() int {
	return len(ph.points)
}

func (ph *pointHeap) Less(i, j int) bool {
	return distance(ph.points[i]) < distance(ph.points[j])
}

func distance(p []int) uint64 {
	x, y := uint64(p[0]), uint64(p[1])
	return x*x + y*y
}

func (ph *pointHeap) Swap(i, j int) {
	ph.points[i], ph.points[j] = ph.points[j], ph.points[i]
}

func (ph *pointHeap) Push(x interface{}) {
	ph.points = append(ph.points, x.([]int))
}

func (ph *pointHeap) Pop() interface{} {
	newLen := len(ph.points) - 1
	x := ph.points[newLen]
	ph.points = ph.points[:newLen]
	return x
}
