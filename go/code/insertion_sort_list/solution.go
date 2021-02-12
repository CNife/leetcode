package insertion_sort_list

import (
	"math"

	. "github.com/CNife/leetcode/go/types"
)

func InsertionSortList(head *ListNode) *ListNode {
	dummyHead := &ListNode{Val: math.MinInt32}
	sort := func(node *ListNode) {
		prev := dummyHead
		for prev.Next != nil {
			if node.Val < prev.Next.Val {
				node.Next, prev.Next = prev.Next, node
				return
			} else {
				prev = prev.Next
			}
		}
		prev.Next, node.Next = node, nil
	}

	for ptr := head; ptr != nil; {
		next := ptr.Next
		sort(ptr)
		ptr = next
	}
	return dummyHead.Next
}
