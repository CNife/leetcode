package minimum_window_substring

func MinWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	pattern := newCounter(t)
	window := newCounter(s[:len(t)])
	i, j := 0, len(t)
	var result string

	for {
		if window.contains(pattern) {
			if i >= j {
				break
			}
			if len(result) < 1 || j-i < len(result) {
				result = s[i:j]
			}
			window.remove(s[i])
			i++
		} else {
			if j >= len(s) {
				break
			}
			window.add(s[j])
			j++
		}
	}

	return result
}

type counter struct {
	inner [52]uint16
}

func newCounter(s string) *counter {
	c := &counter{ /*size: len(s)*/ }
	for i := 0; i < len(s); i++ {
		c.inner[indexOf(s[i])]++
	}
	return c
}

func (c *counter) add(b byte) {
	c.inner[indexOf(b)]++
}

func (c *counter) remove(b byte) {
	p := &c.inner[indexOf(b)]
	if *p > 0 {
		*p--
	}
}

func indexOf(b byte) byte {
	if b < 'a' {
		return b - 'A'
	} else {
		return b - 'a' + 26
	}
}

func (c *counter) contains(other *counter) bool {
	for i := 0; i < 52; i++ {
		if c.inner[i] < other.inner[i] {
			return false
		}
	}
	return true
}
