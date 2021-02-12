package word_break_2

import (
	"reflect"
	"sort"
	"testing"
)

func TestWordBreak(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []*struct {
		s     string
		words []string
		want  []string
	}{
		{
			s:     "catsanddog",
			words: []string{"cat", "cats", "and", "sand", "dog"},
			want:  []string{"cats and dog", "cat sand dog"},
		},
		{
			s:     "pineapplepenapple",
			words: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			want:  []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"},
		},
		{
			s:     "catsandog",
			words: []string{"cats", "dog", "sand", "and", "cat"},
			want:  nil,
		},
		{
			s:     "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			words: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
			want:  nil,
		},
	}
	for _, tt := range tests {
		got, want := sortStrings(WordBreak(tt.s, tt.words)), sortStrings(tt.want)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("WordBreak(%v, %v) = %v, want %v", tt.s, tt.words, got, want)
		}
	}
}

func sortStrings(ss []string) []string {
	sort.Slice(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})
	return ss
}
