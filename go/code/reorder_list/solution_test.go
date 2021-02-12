package reorder_list

import (
	"reflect"
	"testing"

	. "github.com/CNife/leetcode/go/types"
)

func TestReorderList(t *testing.T) {
	tests := []struct {
		head, want *ListNode
	}{
		{NewList(1, 2, 3, 4), NewList(1, 4, 2, 3)},
		{NewList(1, 2, 3, 4, 5), NewList(1, 5, 2, 4, 3)},
	}
	for _, tt := range tests {
		list := tt.head.Clone()
		if ReorderList(list); !reflect.DeepEqual(list, tt.want) {
			t.Errorf("Reorder(%v) = %v, want %v", tt.head, list, tt.want)
		}
	}
}
