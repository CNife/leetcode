package word_pattern

import (
	"strings"
)

func WordPattern(pattern, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	var wordTable [26]string
	wordMap := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		index, word := pattern[i]-'a', words[i]
		foundWord := wordTable[index]
		foundIndex, exists := wordMap[word]
		if !exists && len(foundWord) <= 0 {
			wordTable[index] = word
			wordMap[word] = index
		} else if foundWord != word || foundIndex != index {
			return false
		}
	}
	return true
}
