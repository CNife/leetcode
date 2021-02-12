package word_ladder

import "testing"

func TestLadderLength(t *testing.T) {
	tests := []struct {
		beginWord, endWord string
		wordList           []string
		want               int
	}{
		{
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			want:      5,
		},
		{
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			want:      0,
		},
	}
	for _, tt := range tests {
		if got := LadderLength(tt.beginWord, tt.endWord, tt.wordList); got != tt.want {
			t.Errorf("LadderLength(%v, %v, %v) = %v, want %v", tt.beginWord, tt.endWord, tt.wordList, got, tt.want)
		}
	}
}
