package freedom_trail

import "testing"

func TestFindRotateSteps(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		ring, key string
		want      int
	}{
		{"edcba", "abcde", 10},
		{"godding", "gd", 4},
		{"caotmcaataijjxi", "oatjiioicitatajtijciocjcaaxaaatmctxamacaamjjx", 137},
	}
	for _, tt := range tests {
		if got := FindRotateSteps(tt.ring, tt.key); got != tt.want {
			t.Errorf("FindRotateSteps(%v, %v) = %v, want %v",
				tt.ring, tt.key, got, tt.want)
		}
	}
}
