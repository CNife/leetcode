package populating_next_right_pointers_in_each_node

import (
	"reflect"
	"testing"
)

func TestConnect(t *testing.T) {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val:   2,
			Left:  &Node{Val: 4},
			Right: &Node{Val: 5},
		},
		Right: &Node{
			Val:   3,
			Left:  &Node{Val: 6},
			Right: &Node{Val: 7},
		},
	}
	want := &Node{
		Val: 1,
		Left: &Node{
			Val:   2,
			Left:  &Node{Val: 4},
			Right: &Node{Val: 5},
		},
		Right: &Node{
			Val:   3,
			Left:  &Node{Val: 6},
			Right: &Node{Val: 7},
		},
	}
	want.Left.Next = want.Right
	want.Left.Left.Next = want.Left.Right
	want.Left.Right.Next = want.Right.Left
	want.Right.Left.Next = want.Right.Right

	got := Connect(root)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Connect(%v) = %v, want %v", root, got, want)
	}
}
