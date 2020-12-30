package last_stone_weight

import "container/heap"

func LastStoneWeight(stones []int) int {
	sh := stoneHeap(stones)
	heap.Init(&sh)
	for sh.Len() > 1 {
		m := heap.Pop(&sh).(int)
		n := sh[0]
		if m > n {
			sh[0] = m - n
			heap.Fix(&sh, 0)
		} else {
			heap.Pop(&sh)
		}
	}
	if sh.Len() > 0 {
		return sh[0]
	} else {
		return 0
	}
}

type stoneHeap []int

func (s stoneHeap) Len() int {
	return len(s)
}

func (s stoneHeap) Less(i, j int) bool {
	return s[j] < s[i]
}

func (s stoneHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *stoneHeap) Push(x interface{}) {
	*s = append(*s, x.(int))
}

func (s *stoneHeap) Pop() interface{} {
	old := *s
	newLen := len(old) - 1
	x := old[newLen]
	*s = old[:newLen]
	return x
}
