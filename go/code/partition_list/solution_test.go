package partition_list

import (
	"reflect"
	"testing"

	. "github.com/CNife/leetcode/go/types"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		head *ListNode
		x    int
		want *ListNode
	}{
		{
			head: NewList(1, 4, 3, 2, 5, 2),
			x:    3,
			want: NewList(1, 2, 2, 4, 3, 5),
		},
	}
	for _, tt := range tests {
		if got := Partition(tt.head, tt.x); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Partition(%v, %v) = %v, want %v",
				tt.head, tt.x, got, tt.want)
		}
	}
}
