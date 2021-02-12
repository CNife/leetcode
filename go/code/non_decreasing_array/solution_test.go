package non_decreasing_array

import "testing"

func TestCheckPossibility(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{3, 4, 2, 3},
			want: false,
		},
		{
			nums: []int{4, 2, 3},
			want: true,
		},
		{
			nums: []int{4, 2, 1},
			want: false,
		},
	}
	for _, tt := range tests {
		if got := CheckPossibility(tt.nums); got != tt.want {
			t.Errorf("CheckPossibility(%v) = %v, want %v",
				tt.nums, got, tt.want)
		}
	}
}
