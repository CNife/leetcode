package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewList(t *testing.T) {
	assert.Equal(t, (*ListNode)(nil), NewList())
	assert.Equal(t, &ListNode{Val: 1}, NewList(1))
	assert.Equal(t,
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{Val: 5},
					},
				},
			},
		},
		NewList(1, 2, 3, 4, 5),
	)
}

func TestListNode_String(t *testing.T) {
	assert.Equal(t, "nil", NewList().String())
	assert.Equal(t, "1", NewList(1).String())
	assert.Equal(t, "1->2->3->4->5", NewList(1, 2, 3, 4, 5).String())
	assert.Panics(t, func() {
		_ = loopList().String()
	})
}

func TestListNode_Clone(t *testing.T) {
	assert.Equal(t, (*ListNode)(nil), NewList().Clone())
	assert.Equal(t, NewList(1), NewList(1))
	assert.Equal(t, NewList(1, 2, 3, 4, 5), NewList(1, 2, 3, 4, 5))
	assert.Panics(t, func() {
		_ = loopList().Clone()
	})
}

func loopList() *ListNode {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	list.Next.Next = list

	return list
}
