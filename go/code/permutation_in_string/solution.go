package permutation_in_string

func CheckInclusion(p, s string) bool {
	if len(p) > len(s) {
		return false
	}

	pattern := countString(p)
	window := countString(s[:len(p)])
	i := len(p)
	for {
		if pattern.equals(window) {
			return true
		}
		if i < len(s) {
			window.inc(s[i])
			window.dec(s[i-len(p)])
			i++
		} else {
			return false
		}
	}
}

type counter [26]uint16

func countString(s string) *counter {
	c := &counter{}
	for i := 0; i < len(s); i++ {
		c.inc(s[i])
	}
	return c
}

func (c *counter) inc(b byte) {
	index := b - 'a'
	c[index]++
}

func (c *counter) dec(b byte) {
	index := b - 'a'
	c[index]--
}

func (c *counter) equals(other *counter) bool {
	for i := 0; i < 26; i++ {
		if c[i] != other[i] {
			return false
		}
	}
	return true
}
