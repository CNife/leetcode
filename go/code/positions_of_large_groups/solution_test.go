package positions_of_large_groups

import (
	"reflect"
	"testing"
)

func TestLargeGroupPosition(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s    string
		want [][]int
	}{
		{"abbxxxxzzy", [][]int{{3, 6}}},
		{"abc", nil},
		{"abcdddeeeeaabbbcd", [][]int{{3, 5}, {6, 9}, {12, 14}}},
		{"aba", nil},
		{"aaa", [][]int{{0, 2}}},
	}
	for _, tt := range tests {
		got := LargeGroupPosition(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("LargeGroupPosition(%v) = %v, want %v",
				tt.s, got, tt.want)
		}
	}
}
