package swap_nodes_in_pairs

import . "github.com/CNife/leetcode/go/types"

func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next, newHead.Next = newHead.Next, head

	for prev := head; prev.Next != nil && prev.Next.Next != nil; {
		next, nextNext := prev.Next, prev.Next.Next
		next.Next, nextNext.Next, prev.Next = nextNext.Next, next, nextNext
		prev = next
	}
	return newHead
}
