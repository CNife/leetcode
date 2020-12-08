package substring_with_concatenation_of_all_words

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	tests := []struct {
		s     string
		words []string
		want  []int
	}{
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}, []int{8}},
		{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, nil},
	}
	for _, tt := range tests {
		got := FindSubstring(tt.s, tt.words)
		sort.Ints(got)
		sort.Ints(tt.want)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FindSubstring(%v, %v) = %v, want %v", tt.s, tt.words, got, tt.want)
		}
	}
}
