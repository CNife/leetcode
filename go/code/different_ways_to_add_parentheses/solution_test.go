package different_ways_to_add_parentheses

import (
	"reflect"
	"sort"
	"testing"
)

func TestDiffWaysToCompute(t *testing.T) {
	tests := []struct {
		input string
		want  []int
	}{
		{"2-1-1", []int{0, 2}},
		{"2*3-4*5", []int{-34, -14, -10, -10, 10}},
		{"11", []int{11}},
	}
	for _, tt := range tests {
		got := DiffWaysToCompute(tt.input)
		sort.Ints(got)
		sort.Ints(tt.want)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("DiffWaysToCompute(%v) = %v, want %v",
				tt.input, got, tt.want)
		}
	}
}
