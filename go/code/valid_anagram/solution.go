package valid_anagram

func IsAnagram(s, t string) bool {
	c1, c2 := count(s), count(t)
	return c1 == c2
}

func count(s string) [26]uint {
	var c [26]uint
	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		c[idx]++
	}
	return c
}
