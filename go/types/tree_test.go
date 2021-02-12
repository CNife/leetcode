package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTree(t *testing.T) {
	assert.Equal(t, (*TreeNode)(nil), NewTree())
	assert.Equal(t, &TreeNode{Val: 1}, NewTree(1))
	assert.Equal(t,
		&TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 2},
			Right: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 5},
			},
		},
		NewTree(1, 2, 3, -1, -1, 4, 5))
}

func TestTreeNode_String(t *testing.T) {
	assert.Equal(t, "nil", NewTree().String())
	assert.Equal(t, "(1,nil,nil)", NewTree(1).String())
	assert.Equal(t,
		"(1,(2,nil,nil),(3,(4,nil,nil),(5,nil,nil)))",
		NewTree(1, 2, 3, -1, -1, 4, 5).String())
}

func TestTreeNode_Clone(t *testing.T) {
	assert.Equal(t, NewTree(), NewTree())
	assert.Equal(t, NewTree(1), NewTree(1))
	assert.Equal(t, NewTree(1, 2, 3, -1, 4, 5), NewTree(1, 2, 3, -1, 4, 5))
}
