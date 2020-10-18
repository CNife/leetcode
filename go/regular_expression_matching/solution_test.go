package regular_expression_matching

import "testing"

func TestIsMatch(t *testing.T) {
	tests := []struct {
		s, p string
		want bool
	}{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*p*.", false},
		{"aa", "*a", false},
		{"a", "", false},
	}
	for _, tt := range tests {
		if IsMatch(tt.s, tt.p) != tt.want {
			t.Errorf("IsMatch(%v, %v) = %v, want %v", tt.s, tt.p, !tt.want, tt.want)
		}
	}
}
