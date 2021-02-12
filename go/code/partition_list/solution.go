package partition_list

import . "github.com/CNife/leetcode/go/types"

func Partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var small, big ListNode
	smallPtr, bigPtr := &small, &big
	for head != nil {
		if head.Val < x {
			smallPtr.Next = head
			smallPtr = smallPtr.Next
		} else {
			bigPtr.Next = head
			bigPtr = bigPtr.Next
		}
		head = head.Next
	}

	smallPtr.Next = big.Next
	bigPtr.Next = nil
	return small.Next
}
