package swap_nodes_in_pairs

import (
	"testing"

	. "github.com/CNife/leetcode/go/types"
)

func TestSwapPairs(t *testing.T) {
	tests := []struct {
		head, want *ListNode
	}{
		{NewList(1, 2, 3, 4), NewList(2, 1, 4, 3)},
		{NewList(), NewList()},
		{NewList(1), NewList(1)},
	}
	for _, tt := range tests {
		if got := SwapPairs(tt.head); got.String() != tt.want.String() {
			t.Errorf("SwapPairs(%v) = %v, want %v", tt.head, got, tt.want)
		}
	}
}
