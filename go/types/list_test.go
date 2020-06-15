package types

import (
	"reflect"
	"testing"
)

func TestNewList(t *testing.T) {
	cases := []struct {
		name   string
		values []int
		want   *ListNode
	}{
		{
			name:   "empty list",
			values: []int{},
			want:   nil,
		},
		{
			name:   "non-empty list",
			values: []int{1, 2, 3, 4, 5},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := NewList(c.values...); !reflect.DeepEqual(got, c.want) {
				t.Errorf("NewList(%v) = %v, want %v", c.values, got, c.want)
			}
		})
	}
}
