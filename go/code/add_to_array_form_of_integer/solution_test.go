package add_to_array_form_of_integer

import (
	"reflect"
	"testing"
)

func TestAddToArrayForm(t *testing.T) {
	tests := []struct {
		array []int
		k     int
		want  []int
	}{
		{[]int{1, 2, 0, 0}, 34, []int{1, 2, 3, 4}},
		{[]int{2, 7, 4}, 181, []int{4, 5, 5}},
		{[]int{2, 1, 5}, 806, []int{1, 0, 2, 1}},
		{[]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 1, []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{[]int{6}, 809, []int{8, 1, 5}},
	}
	for _, tt := range tests {
		got := AddToArrayForm(tt.array, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("AddToArrayForm(%v, %v) = %v, want %v",
				tt.array, tt.k, got, tt.want)
		}
	}
}
