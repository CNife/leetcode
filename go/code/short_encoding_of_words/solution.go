package short_encoding_of_words

import "sort"

func MinimumLengthEncoding(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	t := trieNode{}
	result := 0
	for _, word := range words {
		if t.addWordReversed(word) {
			result += len(word) + 1
		}
	}
	return result
}

type trieNode struct {
	children [26]*trieNode
}

func (t *trieNode) addWordReversed(word string) bool {
	node := t
	isNewWord := false
	for i := len(word) - 1; i >= 0; i-- {
		index := word[i] - 'a'
		if node.children[index] == nil {
			node.children[index] = &trieNode{}
			isNewWord = true
		}
		node = node.children[index]
	}
	return isNewWord
}
