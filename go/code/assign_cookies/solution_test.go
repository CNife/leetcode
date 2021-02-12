package assign_cookies

import "testing"

func TestFindContentChildren(t *testing.T) {
	tests := []struct {
		g, s []int
		want int
	}{
		{
			g:    []int{1, 2, 3},
			s:    []int{1, 1},
			want: 1,
		},
		{
			g:    []int{1, 2},
			s:    []int{1, 2, 3},
			want: 2,
		},
	}
	for _, tt := range tests {
		if got := FindContentChildren(tt.g, tt.s); got != tt.want {
			t.Errorf("FindContentChildren(%v, %v) = %v, want %v",
				tt.g, tt.s, got, tt.want)
		}
	}
}
