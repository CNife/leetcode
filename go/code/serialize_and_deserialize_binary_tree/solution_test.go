package serialize_and_deserialize_binary_tree

import (
	"reflect"
	"testing"

	. "github.com/CNife/leetcode/go/types"
)

func TestCodec(t *testing.T) {
	tests := []struct {
		name string
		tree *TreeNode
	}{
		{
			name: "empty tree",
			tree: NewTree(),
		},
		{
			name: "full tree",
			tree: NewTree(1, 2, 3, 4, 5, 6, 7),
		},
		{
			name: "linked list like tree",
			tree: NewTree(1, -1, 2, -1, 3, -1, 4),
		},
	}
	c := Constructor()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := c.deserialize(c.serialize(tt.tree)); !reflect.DeepEqual(actual, tt.tree) {
				t.Errorf("Actual: %v\nExpect: %v", actual, tt.tree)
				t.Fail()
			}
		})
	}
}
