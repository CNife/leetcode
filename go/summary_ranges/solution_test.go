package summary_ranges

import (
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		nums []int
		want []string
	}{
		{[]int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
		{nil, nil},
		{[]int{-1}, []string{"-1"}},
		{[]int{0}, []string{"0"}},
	}
	for _, tt := range tests {
		if got := SummaryRanges(tt.nums); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("SummaryRanges(%v) = %v, want %v",
				tt.nums, got, tt.want)
		}
	}
}
