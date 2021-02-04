package sliding_window_median

import (
	"container/heap"
)

func MedianSlidingWindow(nums []int, k int) []float64 {
	big, small := &minHeap{}, &maxHeap{}
	getMedian := func() float64 {
		if k%2 == 0 {
			return float64(big.top()+small.top()) / 2
		} else {
			return float64(small.top())
		}
	}

	for i := 0; i < k; i++ {
		heap.Push(small, nums[i])
	}
	for i := 0; i < k/2; i++ {
		heap.Push(big, heap.Pop(small))
	}

	remove := map[int]int{}
	result := make([]float64, 0, len(nums)-k+1)
	result = append(result, getMedian())

	for i := k; i < len(nums); i++ {
		balance := 0
		left := nums[i-k]
		if count, exists := remove[left]; exists {
			remove[left] = count + 1
		} else {
			remove[left] = 1
		}

		if small.Len() > 0 && left <= small.top() {
			balance--
		} else {
			balance++
		}
		if small.Len() > 0 && nums[i] <= small.top() {
			heap.Push(small, nums[i])
			balance++
		} else {
			heap.Push(big, nums[i])
			balance--
		}

		if balance > 0 {
			heap.Push(big, heap.Pop(small))
		}
		if balance < 0 {
			heap.Push(small, heap.Pop(big))
		}

		for small.Len() > 0 {
			if count, exists := remove[small.top()]; exists && count > 0 {
				remove[heap.Pop(small).(int)]--
			} else {
				break
			}
		}
		for big.Len() > 0 {
			if count, exists := remove[big.top()]; exists && count > 0 {
				remove[heap.Pop(big).(int)]--
			} else {
				break
			}
		}

		result = append(result, getMedian())
	}

	return result
}

type minHeap struct {
	inner []int
}

func (h *minHeap) top() int {
	return h.inner[0]
}

func (h *minHeap) Len() int {
	return len(h.inner)
}

func (h *minHeap) Less(i, j int) bool {
	return h.inner[i] < h.inner[j]
}

func (h *minHeap) Swap(i, j int) {
	h.inner[i], h.inner[j] = h.inner[j], h.inner[i]
}

func (h *minHeap) Push(x interface{}) {
	h.inner = append(h.inner, x.(int))
}

func (h *minHeap) Pop() interface{} {
	n := len(h.inner)
	x := h.inner[n-1]
	h.inner = h.inner[:n-1]
	return x
}

type maxHeap struct {
	minHeap
}

func (h *maxHeap) Less(i, j int) bool {
	return h.inner[i] > h.inner[j]
}
