package recover_a_tree_from_preorder_traversal

import (
	"reflect"
	"testing"

	. "github.com/CNife/leetcode/go/types"
)

func TestRecoverFromPreorder(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "1",
			args: args{"1-2--3--4-5--6--7"},
			want: NewTree(1, 2, 5, 3, 4, 6, 7),
		},
		{
			name: "2",
			args: args{"1-2--3---4-5--6---7"},
			want: NewTree(1, 2, 5, 3, -1, 6, -1, 4, -1, 7),
		},
		{
			name: "3",
			args: args{"1-401--349---90--88"},
			want: NewTree(1, 401, -1, 349, 88, 90),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RecoverFromPreorder(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RecoverFromPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
