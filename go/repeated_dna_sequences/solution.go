package repeated_dna_sequences

func FindRepeatedDnaSequences(src string) []string {
	if len(src) < 10 {
		return nil
	}

	var ptn pattern = 0
	for i := 0; i < 10; i++ {
		ptn = feed(ptn, src[i])
	}

	knownPatterns := map[pattern]int{
		ptn: 1,
	}
	for i := 10; i < len(src); i++ {
		ptn = slide(ptn, src[i])
		if count, exists := knownPatterns[ptn]; exists {
			knownPatterns[ptn] = count + 1
		} else {
			knownPatterns[ptn] = 1
		}
	}

	var result []string
	for p, c := range knownPatterns {
		if c > 1 {
			result = append(result, p.String())
		}
	}
	return result
}

type pattern uint32

func slide(ptn pattern, ch byte) pattern {
	return feed(ptn&0x3ffff, ch)
}

func feed(ptn pattern, ch byte) pattern {
	return (ptn << 2) | char2Bits(ch)
}

func char2Bits(b byte) pattern {
	switch b {
	case 'A':
		return 0
	case 'C':
		return 1
	case 'G':
		return 2
	case 'T':
		return 3
	default:
		panic("invalid char")
	}
}

func (p pattern) String() string {
	result := make([]byte, 10)
	for i := 9; i >= 0; i-- {
		switch p & 0b11 {
		case 0:
			result[i] = 'A'
		case 1:
			result[i] = 'C'
		case 2:
			result[i] = 'G'
		case 3:
			result[i] = 'T'
		}
		p >>= 2
	}
	return string(result)
}
