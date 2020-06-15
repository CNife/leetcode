package types

import (
	"reflect"
	"testing"
)

func TestNewTree(t *testing.T) {
	cases := []struct {
		name   string
		values []int
		expect *TreeNode
	}{
		{
			name:   "empty tree",
			values: []int{},
			expect: nil,
		},
		{
			name:   "non-empty tree",
			values: []int{1, 2, 3, 4, 5},
			expect: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{Val: 3},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if actual := NewTree(c.values...); !reflect.DeepEqual(actual, c.expect) {
				t.Errorf("NewTree(%v) = %v\nwant %v", c.values, actual, c.expect)
			}
		})
	}
}
