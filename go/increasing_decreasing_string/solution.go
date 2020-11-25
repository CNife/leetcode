package increasing_decreasing_string

import "bytes"

func SortString(s string) string {
	var buckets [26]uint
	for i := 0; i < len(s); i++ {
		buckets[s[i]-'a']++
	}

	var buffer bytes.Buffer
	for {
		finished := true
		for i := 0; i < 26; i++ {
			if buckets[i] > 0 {
				finished = false
				buckets[i]--
				buffer.WriteByte(byte(i + 'a'))
			}
		}
		for i := 25; i > -1; i-- {
			if buckets[i] > 0 {
				finished = false
				buckets[i]--
				buffer.WriteByte(byte(i + 'a'))
			}
		}
		if finished {
			break
		}
	}
	return buffer.String()
}
