package unique_binary_search_trees_2

import (
	"reflect"
	"testing"

	. "github.com/CNife/leetcode/go/types"
)

func TestGenerateTrees(t *testing.T) {
	tests := []struct {
		n    int
		want []*TreeNode
	}{
		//{
		//	n: 2,
		//	want: []*TreeNode{
		//		NewTree(1, -1, 2),
		//		NewTree(2, 1),
		//	},
		//},
		//{
		//	n: 1,
		//	want: []*TreeNode{
		//		NewTree(1),
		//	},
		//},
		{
			n: 3,
			want: []*TreeNode{
				NewTree(1, -1, 3, 2),
				NewTree(3, 2, -1, 1),
				NewTree(3, 1, -1, -1, 2),
				NewTree(2, 1, 3),
				NewTree(1, -1, 2, -1, 3),
			},
		},
		{
			n:    0,
			want: nil,
		},
	}

	slice2Set := func(trees []*TreeNode) map[string]struct{} {
		set := make(map[string]struct{}, len(trees))
		for _, tree := range trees {
			serialized := tree.String()
			if _, exists := set[serialized]; exists {
				t.Errorf("duplicated %v", trees)
			}
			set[serialized] = struct{}{}
		}
		return set
	}

	for _, tt := range tests {
		got := GenerateTrees(tt.n)
		gotSet, wantSet := slice2Set(got), slice2Set(tt.want)
		if !reflect.DeepEqual(gotSet, wantSet) {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}
