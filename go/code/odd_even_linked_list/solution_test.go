package odd_even_linked_list

import (
	"reflect"
	"testing"

	"github.com/CNife/leetcode/go/types"
)

func TestOddEvenList(t *testing.T) {
	tests := []struct {
		head, want *types.ListNode
	}{
		{types.NewList(1, 2, 3, 4, 5), types.NewList(1, 3, 5, 2, 4)},
		{types.NewList(2, 1, 3, 5, 6, 4, 7), types.NewList(2, 3, 6, 7, 1, 5, 4)},
	}
	for _, tt := range tests {
		if got := OddEvenList(tt.head); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("OddEvenList(%v) = %v, want %v", tt.head, got, tt.want)
		}
	}
}
