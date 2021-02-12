package binary_tree_preorder_traversal

import (
	"reflect"
	"testing"

	. "github.com/CNife/leetcode/go/types"
)

func TestPreorderTraversal(t *testing.T) {
	tests := []*TreeNode{
		nil,
		NewTree(1, -1, 2, 3),
	}
	for _, tt := range tests {
		want := PreorderTraversalRecursion(tt)
		got := PreorderTraversalIteration(tt)
		if !reflect.DeepEqual(want, got) {
			t.Errorf("PreorderTraversal(%v) = %v, want %v", tt, got, want)
		}
	}
}
