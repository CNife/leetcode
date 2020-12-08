package substring_with_concatenation_of_all_words

func FindSubstring(s string, words []string) []int {
	if len(words) <= 0 {
		return nil
	}
	wordLen := len(words[0])
	windowWidth := wordLen * len(words)
	if len(s) < windowWidth {
		return nil
	}

	var want counter
	for _, word := range words {
		want.eat(word)
	}

	var window counter
	window.eat(s[:windowWidth])
	ws := initWordSearcher(words)
	var result []int
	for end := windowWidth; end <= len(s); end++ {
		start := end - windowWidth
		if window.match(&want) && ws.match(s[start:end]) {
			result = append(result, start)
		}
		if end < len(s) {
			window.increase(s[end])
		}
		window.decrease(s[start])
	}
	return result
}

type counter [26]int

func (c *counter) eat(s string) {
	for i := 0; i < len(s); i++ {
		c[s[i]-'a']++
	}
}

func (c *counter) increase(b byte) {
	c[b-'a']++
}

func (c *counter) decrease(b byte) {
	c[b-'a']--
}

func (c *counter) match(other *counter) bool {
	for i := 0; i < 26; i++ {
		if c[i] != other[i] {
			return false
		}
	}
	return true
}

type wordSearcher struct {
	counter map[string]int
	wordLen int
}

func initWordSearcher(words []string) wordSearcher {
	counter := make(map[string]int)
	for _, word := range words {
		if count, exists := counter[word]; exists {
			counter[word] = count + 1
		} else {
			counter[word] = 1
		}
	}
	return wordSearcher{
		counter: counter,
		wordLen: len(words[0]),
	}
}

func (ws wordSearcher) match(s string) bool {
	other := make(map[string]int)
	for word, count := range ws.counter {
		other[word] = count
	}

	for i := 0; i < len(s)/ws.wordLen; i++ {
		word := s[i*ws.wordLen : (i+1)*ws.wordLen]
		if count, exists := other[word]; exists {
			if count <= 1 {
				delete(other, word)
			} else {
				other[word] = count - 1
			}
		} else {
			return false
		}
	}
	return true
}
