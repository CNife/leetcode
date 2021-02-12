package remove_nth_node_from_end_of_list

import (
	"testing"

	. "github.com/CNife/leetcode/go/types"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		head *ListNode
		n    int
		want *ListNode
	}{
		{
			head: NewList(1, 2, 3, 4, 5),
			n:    2,
			want: NewList(1, 2, 3, 5),
		},
	}
	for _, tt := range tests {
		if got := RemoveNthFromEnd(tt.head, tt.n); got.String() != tt.want.String() {
			t.Errorf("RemoveNthFromEnd(%v, %v) = %v, want %v", tt.head, tt.n, got, tt.want)
		}
	}
}
