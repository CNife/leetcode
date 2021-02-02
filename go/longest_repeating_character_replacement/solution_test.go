package longest_repeating_character_replacement

import "testing"

func TestCharacterReplacement(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s       string
		k, want int
	}{
		{
			s:    "ABAB",
			k:    2,
			want: 4,
		},
		{
			s:    "AABABBA",
			k:    1,
			want: 4,
		},
	}
	for _, tt := range tests {
		if got := CharacterReplacement(tt.s, tt.k); got != tt.want {
			t.Errorf("CharacterReplacement(%v, %v) = %v, want %v",
				tt.s, tt.k, got, tt.want)
		}
	}
}
