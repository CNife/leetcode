package remove_duplicate_letters

import "testing"

func TestRemoveDuplicateLetters(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s, want string
	}{
		{"bcabc", "abc"},
		{"cbacdcbc", "acdb"},
	}
	for _, tt := range tests {
		if got := RemoveDuplicateLetters(tt.s); got != tt.want {
			t.Errorf("RemoveDuplicateLetters(%v) = %v, want %v",
				tt.s, got, tt.want)
		}
	}
}
