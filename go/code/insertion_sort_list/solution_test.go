package insertion_sort_list

import (
	"reflect"
	"testing"

	. "github.com/CNife/leetcode/go/types"
)

func TestInsertionSortList(t *testing.T) {
	tests := []struct {
		head, want *ListNode
	}{
		{
			head: nil,
			want: nil,
		},
		{
			head: NewList(1),
			want: NewList(1),
		},
		{
			head: NewList(4, 2, 1, 3),
			want: NewList(1, 2, 3, 4),
		},
		{
			head: NewList(-1, 5, 3, 4, 0),
			want: NewList(-1, 0, 3, 4, 5),
		},
	}
	for _, tt := range tests {
		if got := InsertionSortList(tt.head); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("InsertionSortList(%v) = %v, want %v",
				tt.head, got, tt.want)
		}
	}
}
