package types

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(values ...int) *ListNode {
	var head, prev *ListNode
	for _, value := range values {
		next := &ListNode{Val: value}
		if prev == nil {
			prev = next
			head = next
		} else {
			prev.Next = next
			prev = next
		}
	}
	return head
}
