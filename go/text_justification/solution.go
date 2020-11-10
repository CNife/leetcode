package text_justification

import (
	"bytes"
)

func FullJustify(words []string, maxWidth int) []string {
	lines := distributeWords(words, maxWidth)
	var result []string
	for i := 0; i < len(lines)-1; i++ {
		if len(lines[i]) == 1 {
			result = append(result, singleWord(lines[i][0], maxWidth))
		} else {
			result = append(result, connectLine(lines[i], maxWidth))
		}
	}
	if len(lines) > 0 {
		lastLine := connectLastLine(lines[len(lines)-1], maxWidth)
		result = append(result, lastLine)
	}
	return result
}

func distributeWords(words []string, maxWidth int) [][]string {
	var result [][]string
	line := []string{words[0]}
	width := len(words[0])
	for i := 1; i < len(words); i++ {
		word, wordLen := words[i], len(words[i])
		if width+wordLen+1 <= maxWidth {
			line = append(line, word)
			width += wordLen + 1
		} else {
			result = append(result, line)
			line = []string{word}
			width = wordLen
		}
	}
	return append(result, line)
}

func singleWord(word string, maxWidth int) string {
	width := len(word)
	if width < maxWidth {
		var buf bytes.Buffer
		buf.Grow(maxWidth)
		buf.WriteString(word)
		for i := 0; i < maxWidth-width; i++ {
			buf.WriteByte(' ')
		}
		return buf.String()
	}
	return word
}

func connectLine(words []string, maxWidth int) string {
	spaceCount := maxWidth - wordsLength(words)
	baseSpaces, moreSpaceCount := spaceCount/(len(words)-1), spaceCount%(len(words)-1)

	var buf bytes.Buffer
	buf.Grow(maxWidth)
	buf.WriteString(words[0])
	for i := 1; i < len(words); i++ {
		for j := 0; j < baseSpaces; j++ {
			buf.WriteByte(' ')
		}
		if moreSpaceCount > 0 {
			buf.WriteByte(' ')
			moreSpaceCount--
		}
		buf.WriteString(words[i])
	}
	return buf.String()
}

func wordsLength(words []string) int {
	result := 0
	for _, word := range words {
		result += len(word)
	}
	return result
}

func connectLastLine(words []string, maxWidth int) string {
	var buf bytes.Buffer
	buf.Grow(maxWidth)
	buf.WriteString(words[0])
	width := len(words[0])
	for i := 1; i < len(words); i++ {
		buf.WriteByte(' ')
		buf.WriteString(words[i])
		width += len(words[i]) + 1
	}
	for j := 0; j < maxWidth-width; j++ {
		buf.WriteByte(' ')
	}
	return buf.String()
}
