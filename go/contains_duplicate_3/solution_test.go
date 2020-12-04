package contains_duplicate_3

import "testing"

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		k, t int
		want bool
	}{
		{[]int{2147483647, -1, 2147483647}, 1, 2147483647, false},
		{[]int{1, 3, 6, 2}, 1, 2, true},
		{[]int{0, 10, 22, 15, 0, 5, 22, 12, 1, 5}, 3, 3, false},
		{[]int{1, 2, 3, 1}, 3, 0, true},
		{[]int{1, 0, 1, 1}, 1, 2, true},
		{[]int{1, 5, 9, 1, 5, 9}, 2, 3, false},
	}
	for _, tt := range tests {
		got := ContainsNearbyAlmostDuplicate(tt.nums, tt.k, tt.t)
		if got != tt.want {
			t.Errorf("ContainsNearbyAlmostDuplicate(%v, %v, %v) = %v, want %v",
				tt.nums, tt.k, tt.t, got, tt.want)
		}
	}
}
