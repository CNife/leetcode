package flatten_binary_tree_to_linked_list

import (
	"leetcode/types"
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	tests := []struct {
		root *types.TreeNode
		want *types.TreeNode
	}{
		{
			root: types.NewTree(1, 2, 5, 3, 4, -1, 6),
			want: types.NewTree(1, -1, 2, -1, 3, -1, 4, -1, 5, -1, 6),
		},
	}
	for _, tt := range tests {
		got := tt.root.DeepClone()
		Flatten(got)
		if !reflect.DeepEqual(got, tt.want) {
			t.Error(tt, got)
		}
	}
}
