package first_unique_character_in_a_string

import "math"

func FirstUnique(s string) int {
	var isMulti [26]bool
	var index [26]int
	for i := 0; i < 26; i++ {
		index[i] = -1
	}

	for i := 0; i < len(s); i++ {
		ch := s[i] - 'a'
		if isMulti[ch] {
			continue
		}
		if index[ch] < 0 {
			index[ch] = i
		} else {
			isMulti[ch] = true
		}
	}

	result := math.MaxInt32
	for i := 0; i < 26; i++ {
		if !isMulti[i] && index[i] >= 0 && index[i] < result {
			result = index[i]
		}
	}
	if result == math.MaxInt32 {
		return -1
	} else {
		return result
	}
}
