package find_the_difference

func FindTheDifference(s, t string) byte {
	var byteTable [26]int
	for i := 0; i < len(s); i++ {
		byteTable[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		byteTable[t[i]-'a']--
		if byteTable[t[i]-'a'] < 0 {
			return t[i]
		}
	}
	return 0
}
