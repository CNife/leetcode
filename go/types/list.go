package types

import (
	"bytes"
	"fmt"
	"strconv"
)

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

func (l *ListNode) String() string {
	if l == nil {
		return "nil"
	}

	nodes := make(map[*ListNode]struct{})
	buffer := bytes.NewBuffer([]byte(strconv.Itoa(l.Val)))
	for ptr := l.Next; ptr != nil; ptr = ptr.Next {
		if _, exists := nodes[ptr]; exists {
			buffer.WriteString(fmt.Sprintf(" <recursion at %v>", ptr))
			break
		} else {
			buffer.WriteString("->")
			buffer.WriteString(strconv.Itoa(ptr.Val))
		}
	}
	return buffer.String()
}

func (l *ListNode) Clone() *ListNode {
	if l == nil {
		return nil
	}
	return &ListNode{
		Val:  l.Val,
		Next: l.Next.Clone(),
	}
}
