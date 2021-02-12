package subarrays_with_k_different_integers

import "testing"

func TestSubarraysWithKDistinct(t *testing.T) {
	tests := []struct {
		nums    []int
		k, want int
	}{
		{
			nums: []int{1, 2, 1, 2, 3}, k: 2,
			want: 7,
		},
		{
			nums: []int{1, 2, 1, 3, 4}, k: 3,
			want: 3,
		},
	}
	for _, tt := range tests {
		got := SubarraysWithKDistinct(tt.nums, tt.k)
		if got != tt.want {
			t.Errorf("SubarraysWithKDistinct(%v, %v) = %v, want %v",
				tt.nums, tt.k, got, tt.want)
		}
	}
}
