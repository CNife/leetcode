package longest_continuous_increasing_subsequence

import "testing"

func TestFindLengthOfLCIS(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 3, 5, 4, 7},
			want: 3,
		},
		{
			nums: []int{2, 2, 2, 2, 2},
			want: 1,
		},
	}
	for _, tt := range tests {
		if got := FindLengthOfLCIS(tt.nums); got != tt.want {
			t.Errorf("FindLengthOfLCIS(%v) = %v, want %v",
				tt.nums, got, tt.want)
		}
	}
}
