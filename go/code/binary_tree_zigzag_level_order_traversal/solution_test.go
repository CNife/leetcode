package binary_tree_zigzag_level_order_traversal

import (
	"reflect"
	"testing"

	. "github.com/CNife/leetcode/go/types"
)

func TestZigzagLevelOrder(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want [][]int
	}{
		{
			root: NewTree(3, 9, 20, -1, -1, 15, 7),
			want: [][]int{{3}, {20, 9}, {15, 7}},
		},
	}
	for _, tt := range tests {
		got := ZigzagLevelOrder(tt.root)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("ZigzagLevelOrder(%v) = %v, want %v",
				tt.root, got, tt.want)
		}
	}
}
