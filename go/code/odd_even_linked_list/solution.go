package odd_even_linked_list

import . "github.com/CNife/leetcode/go/types"

func OddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	flag := true
	evenHead, ptr := head.Next, head.Next.Next
	odds, evens := head, head.Next
	for ptr != nil {
		if flag {
			odds.Next = ptr
			odds = ptr
		} else {
			evens.Next = ptr
			evens = ptr
		}
		ptr = ptr.Next
		flag = !flag
	}
	odds.Next = evenHead
	evens.Next = nil
	return head
}
