package text_justification

import (
	"reflect"
	"testing"
)

func TestFullJustify(t *testing.T) {
	tests := []struct {
		words    []string
		maxWidth int
		want     []string
	}{
		{
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
			want:     []string{"This    is    an", "example  of text", "justification.  "},
		},
		{
			words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
			want:     []string{"What   must   be", "acknowledgment  ", "shall be        "},
		},
		{
			words: []string{
				"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a",
				"computer.", "Art", "is", "everything", "else", "we", "do",
			},
			maxWidth: 20,
			want: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
	}
	for _, tt := range tests {
		if got := FullJustify(tt.words, tt.maxWidth); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FullJustify(%v, %v) = %v, want %v",
				tt.words, tt.maxWidth, got, tt.want)
		}
	}
}
