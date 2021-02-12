package types

import (
	"strconv"
	"strings"
)

// 单链表长度上限
const SizeThreshold = 1000

// 单链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 创建新的单链表
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

// 打印单链表
// 如果链表有环，将引发错误
func (l *ListNode) String() string {
	if l == nil {
		return "nil"
	}

	ptr := l
	var buffer strings.Builder
	buffer.WriteString(strconv.Itoa(ptr.Val))
	ptr = ptr.Next

	var size int
	for ptr != nil {
		if size > SizeThreshold {
			panic("list too long or has loop")
		}

		buffer.WriteString("->")
		buffer.WriteString(strconv.Itoa(ptr.Val))

		ptr = ptr.Next
		size++
	}

	return buffer.String()
}

// 克隆单链表
// 如果链表有环，将引发错误
func (l *ListNode) Clone() *ListNode {
	if l == nil {
		return nil
	}

	head := &ListNode{Val: l.Val}
	oldPtr, newPtr := l.Next, head
	size := 0
	for oldPtr != nil {
		if size > SizeThreshold {
			panic("list too long or has loop")
		}

		newPtr.Next = &ListNode{Val: oldPtr.Val}
		oldPtr, newPtr = oldPtr.Next, newPtr.Next
		size++
	}
	return head
}
