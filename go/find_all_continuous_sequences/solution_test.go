package find_all_continuous_sequences

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findContinuousSequence(t *testing.T) {
	type args struct {
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{9},
			want: [][]int{
				{2, 3, 4},
				{4, 5},
			},
		},
		{
			name: "2",
			args: args{15},
			want: [][]int{
				{1, 2, 3, 4, 5},
				{4, 5, 6},
				{7, 8},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Slice(tt.want, sortFunc(tt.want))
			got := findContinuousSequence(tt.args.target)
			sort.Slice(got, sortFunc(got))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findContinuousSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sortFunc(slice [][]int) func(i, j int) bool {
	return func(i, j int) bool {
		lhs, rhs := slice[i], slice[j]
		if len(lhs) > len(rhs) {
			lhs, rhs = rhs, lhs
		}
		for i := 0; i < len(lhs); i++ {
			if lhs[i] != rhs[i] {
				return lhs[i] < rhs[i]
			}
		}
		return len(lhs) < len(rhs)
	}
}
