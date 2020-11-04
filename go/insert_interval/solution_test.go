package insert_interval

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			want:        [][]int{{1, 5}, {6, 9}},
		},
		{
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			want:        [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
	}
	for _, tt := range tests {
		if got := Insert(tt.intervals, tt.newInterval); !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("Insert(%v, %v) = %v, want %v", tt.intervals, tt.newInterval, got, tt.want)
		}
	}
}
