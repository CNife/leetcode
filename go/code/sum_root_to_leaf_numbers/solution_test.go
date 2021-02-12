package sum_root_to_leaf_numbers

import (
	"testing"

	. "github.com/CNife/leetcode/go/types"
)

func TestSumNumbers(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{NewTree(1, 2, 3), 25},
		{NewTree(4, 9, 0, 5, 1), 1026},
	}
	for _, tt := range tests {
		if got := SumNumbers(tt.root); got != tt.want {
			t.Errorf("SumNumbers(%v) = %v, want %v", tt.root, got, tt.want)
		}
	}
}
