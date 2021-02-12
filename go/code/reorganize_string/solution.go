package reorganize_string

import (
	"bytes"
	"container/heap"
)

func ReorganizeString(src string) string {
	var items [26]int
	for i := 0; i < len(src); i++ {
		index := src[i] - 'a'
		items[index]++
	}

	c := &charHeap{}
	for i := 0; i < 26; i++ {
		if items[i] > 0 {
			item := charHeapItem{
				char:  byte(i) + 'a',
				count: items[i],
			}
			c.inner = append(c.inner, item)
		}
	}
	heap.Init(c)

	var buf bytes.Buffer
	for c.Len() > 1 {
		buf.WriteByte(c.inner[0].char)
		buf.WriteByte(c.inner[1].char)

		c.inner[1].count--
		if c.inner[1].count < 1 {
			heap.Remove(c, 1)
		} else {
			heap.Fix(c, 1)
		}

		c.inner[0].count--
		if c.inner[0].count < 1 {
			heap.Pop(c)
		} else {
			heap.Fix(c, 0)
		}
	}

	if c.Len() == 1 {
		if c.inner[0].count == 1 {
			buf.WriteByte(c.inner[0].char)
			return buf.String()
		} else {
			return ""
		}
	}
	return buf.String()
}

type charHeap struct {
	inner []charHeapItem
}

type charHeapItem struct {
	char  byte
	count int
}

func (c *charHeap) Len() int {
	return len(c.inner)
}

func (c *charHeap) Less(i, j int) bool {
	return c.inner[j].count < c.inner[i].count
}

func (c *charHeap) Swap(i, j int) {
	c.inner[i], c.inner[j] = c.inner[j], c.inner[i]
}

func (c *charHeap) Push(x interface{}) {
	c.inner = append(c.inner, x.(charHeapItem))
}

func (c *charHeap) Pop() interface{} {
	n := len(c.inner)
	x := c.inner[n-1]
	c.inner = c.inner[:n-1]
	return x
}
