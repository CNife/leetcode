package get_equal_substrings_within_budget

import "testing"

func TestEqualSubstring(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s, t          string
		maxCost, want int
	}{
		{
			s: "abcd", t: "bcdf", maxCost: 3,
			want: 3,
		},
		{
			s: "abcd", t: "cdef", maxCost: 3,
			want: 1,
		},
		{
			s: "abcd", t: "acde", maxCost: 0,
			want: 1,
		},
	}
	for _, tt := range tests {
		got := EqualSubstring(tt.s, tt.t, tt.maxCost)
		if got != tt.want {
			t.Errorf("EqualSubstring(%v, %v, %v) = %v, want %v",
				tt.s, tt.t, tt.maxCost, got, tt.want)
		}
	}
}
