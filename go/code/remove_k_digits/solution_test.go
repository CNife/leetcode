package remove_k_digits

import "testing"

func TestRemoveKDigits(t *testing.T) {
	tests := []struct {
		num  string
		k    int
		want string
	}{
		{"1432219", 3, "1219"},
		{"10200", 1, "200"},
		{"10", 1, "0"},
	}
	for _, tt := range tests {
		if got := RemoveKDigits(tt.num, tt.k); got != tt.want {
			t.Errorf("RemoveKDigits(%v, %v) = %v, want %v",
				tt.num, tt.k, got, tt.want)
		}
	}
}
