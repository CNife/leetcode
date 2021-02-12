package longest_repeating_character_replacement

func CharacterReplacement(s string, k int) int {
	var (
		left, right int
		charMax     int
		counter     [26]int
	)

	for ; right < len(s); right++ {
		index := s[right] - 'A'
		counter[index]++
		charMax = max(charMax, counter[index])
		if right-left+1 > charMax+k {
			counter[s[left]-'A']--
			left++
		}
	}

	return len(s) - left
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
