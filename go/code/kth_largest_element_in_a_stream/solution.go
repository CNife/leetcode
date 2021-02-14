package kth_largest_element_in_a_stream

import "container/heap"

type KthLargest struct {
	k    int
	heap minHeap
}

func Constructor(k int, nums []int) KthLargest {
	if k <= 0 {
		panic("invalid k")
	}

	result := KthLargest{k: k}
	if len(nums) > k {
		result.heap = nums[:k]
	} else {
		result.heap = nums
	}

	heap.Init(&result.heap)

	if len(nums) > k {
		for i := k; i < len(nums); i++ {
			_ = result.Add(nums[i])
		}
	}

	return result
}

func (l *KthLargest) Add(val int) int {
	// 0 <= len(l.heap) <= l.k
	switch len(l.heap) {
	case l.k:
		if val > l.heap[0] {
			l.heap[0] = val
			heap.Fix(&l.heap, 0)
		}
		return l.heap[0]
	case l.k - 1:
		heap.Push(&l.heap, val)
		return l.heap[0]
	case 0:
		return -1
	default:
		heap.Push(&l.heap, val)
		return -1
	}
}

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
