package add_binary

import "strings"

func addBinary(a, b string) string {
	var ad adder
	i, j := len(a)-1, len(b)-1
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		ad.addDigit(a[i], b[j])
	}
	for ; i >= 0; i-- {
		ad.addDigit(a[i], '0')
	}
	for ; j >= 0; j-- {
		ad.addDigit('0', b[j])
	}
	return ad.string()
}

type adder struct {
	buffer strings.Builder
	carry  bool
}

func (a *adder) addDigit(lhs, rhs byte) {
	aDigit, bDigit := lhs-'0', rhs-'0'
	result := aDigit + bDigit
	if a.carry {
		result++
	}
	if result > 1 {
		a.buffer.WriteByte(result - 2 + '0')
		a.carry = true
	} else {
		a.buffer.WriteByte(result + '0')
		a.carry = false
	}
}

func (a *adder) string() string {
	if a.carry {
		a.buffer.WriteByte('1')
	}
	s := []byte(a.buffer.String())
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}
