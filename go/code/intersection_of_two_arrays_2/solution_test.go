package intersection_of_two_arrays_2

import (
	"reflect"
	"testing"
)

func Test_intersect(t *testing.T) {
	type args struct {
		lhs []int
		rhs []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{[]int{1, 2, 2, 1}, []int{2, 2}},
			want: []int{2, 2},
		},
		{
			name: "2",
			args: args{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}},
			want: []int{4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildMap(intersect(tt.args.lhs, tt.args.rhs))
			want := buildMap(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
