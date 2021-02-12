package sort_list

import (
	"reflect"
	"testing"

	. "github.com/CNife/leetcode/go/types"
)

func TestSortList(t *testing.T) {
	tests := []struct {
		head, want *ListNode
	}{
		{
			head: NewList(4, 2, 1, 3),
			want: NewList(1, 2, 3, 4),
		},
		{
			head: NewList(-1, 5, 3, 4, 0),
			want: NewList(-1, 0, 3, 4, 5),
		},
		{
			head: nil,
			want: nil,
		},
		{
			head: NewList(0),
			want: NewList(0),
		},
	}
	for _, tt := range tests {
		if got := SortList(tt.head); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("SortList(%v) = %v, want %v", tt.head, got, tt.want)
		}
	}
}
