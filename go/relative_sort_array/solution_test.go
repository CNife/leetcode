package relative_sort_array

import (
	"reflect"
	"testing"
)

func TestRelativeSortArray(t *testing.T) {
	tests := []struct {
		arr1, arr2, want []int
	}{
		{
			arr1: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			arr2: []int{2, 1, 4, 3, 9, 6},
			want: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
	}
	for _, tt := range tests {
		got := RelativeSortArray(tt.arr1, tt.arr2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("RelativeSortArray(%v, %v) = %v, want %v",
				tt.arr1, tt.arr2, got, tt.want)
		}
	}
}
