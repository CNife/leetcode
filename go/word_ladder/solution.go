package word_ladder

func LadderLength(beginWord, endWord string, wordList []string) int {
	words := newStringSet(wordList...)
	if !words.contains(endWord) {
		return 0
	}
	beginSet, endSet := newStringSet(beginWord), newStringSet(endWord)
	visited := newStringSet()

	for length := 1; beginSet.hasValues() && endSet.hasValues(); length++ {
		if beginSet.size() > endSet.size() {
			beginSet, endSet = endSet, beginSet
		}
		tmpSet := newStringSet()
		for word := range beginSet {
			buffer := []byte(word)
			for i, oldByte := range buffer {
				for replaceByte := byte('a'); replaceByte <= 'z'; replaceByte++ {
					if oldByte == replaceByte {
						continue
					}
					buffer[i] = replaceByte
					tmpWord := string(buffer)
					if endSet.contains(tmpWord) {
						return length + 1
					}
					if !visited.contains(tmpWord) && words.contains(tmpWord) {
						visited.insert(tmpWord)
						tmpSet.insert(tmpWord)
					}
				}
				buffer[i] = oldByte
			}
		}
		beginSet = tmpSet
	}
	return 0
}

type stringSet map[string]struct{}

func newStringSet(words ...string) stringSet {
	result := make(stringSet)
	for _, word := range words {
		result[word] = struct{}{}
	}
	return result
}

func (s stringSet) contains(word string) bool {
	_, exists := s[word]
	return exists
}

func (s stringSet) hasValues() bool {
	return len(s) > 0
}

func (s stringSet) size() int {
	return len(s)
}

func (s stringSet) insert(word string) {
	s[word] = struct{}{}
}
